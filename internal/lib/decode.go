package lib

import (
	"bytes"
	"io"
	"sync"
	"unsafe"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
	"github.com/Tnze/go-mc/nbt"
	"github.com/klauspost/compress/gzip"
)

var gzipReaderPool = sync.Pool{
	New: func() any {
		return &gzip.Reader{}
	},
}

// Pool for byte buffers to reduce allocations
// Use 32KB buffer for better performance with larger inventories
var bufferPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 32768) // 32KB pre-allocation
		return &b
	},
}

// Inflated NBT, so the decoder reads from memory.
var inflatePool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

var bytesReaderPool = sync.Pool{
	New: func() any {
		return &bytes.Reader{}
	},
}

// stringToBytes converts string to []byte without allocation using unsafe
// UNSAFE: The returned slice MUST NOT be modified
//
//go:inline
func stringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func DecodeInventory(data string) (*models.DecodedInventory, error) {
	if data == "" {
		return &models.DecodedInventory{}, nil
	}

	// Pre-calculate decoded length
	decodedLen := DecodedLen(len(data))

	// Get buffer from pool
	bufPtr := bufferPool.Get().(*[]byte)
	buf := *bufPtr

	// Ensure buffer is large enough
	if cap(buf) < decodedLen {
		buf = make([]byte, decodedLen)
		*bufPtr = buf
	} else {
		buf = buf[:decodedLen:decodedLen]
	}

	// Decode base64 using unsafe string-to-bytes (zero-copy)
	n, err := FastDecode(buf, stringToBytes(data))
	if err != nil {
		bufferPool.Put(bufPtr)
		return nil, err
	}

	// Decode the gzip + NBT data
	result, err := DecodeFromBytes(buf[:n])
	bufferPool.Put(bufPtr)
	return result, err
}

func DecodeFromBytes(data []byte) (*models.DecodedInventory, error) {
	reader := gzipReaderPool.Get().(*gzip.Reader)
	bytesReader := bytesReaderPool.Get().(*bytes.Reader)

	bytesReader.Reset(data)
	err := reader.Reset(bytesReader)
	if err != nil {
		gzipReaderPool.Put(reader)
		bytesReaderPool.Put(bytesReader)
		return nil, err
	}

	// Inflate first: the NBT decoder reads byte by byte, and a gzip reader has no ReadByte, so decoding straight
	// from it cost a gzip Read call per byte (about 20% slower on real inventories).
	inflated := inflatePool.Get().(*bytes.Buffer)
	inflated.Reset()
	_, err = io.Copy(inflated, reader)

	// Close before pooling: once Put, another goroutine may already be reusing the reader.
	_ = reader.Close()
	gzipReaderPool.Put(reader)
	if err != nil {
		inflatePool.Put(inflated)
		bytesReaderPool.Put(bytesReader)
		return nil, err
	}

	// Not pooled: the result goes to the caller and is never returned, and a pooled 256-slot slice per call
	// cost far more than letting the decoder size the slice to the real item count.
	nbtData := &models.DecodedInventory{}

	// The decoder copies every string and array it reads, so nothing it returns points into the pooled buffer.
	bytesReader.Reset(inflated.Bytes())
	_, err = nbt.NewDecoder(bytesReader).Decode(nbtData)
	inflatePool.Put(inflated)
	bytesReaderPool.Put(bytesReader)

	if err != nil {
		return nil, err
	}

	return nbtData, nil
}
