package lib

import (
	"bytes"
	"sync"
	"unsafe"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
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

var bytesReaderPool = sync.Pool{
	New: func() any {
		return &bytes.Reader{}
	},
}

// Pool for decoded inventory structs to reduce allocations
// Pre-allocate larger capacity to handle most inventories without reallocation
var decodedInventoryPool = sync.Pool{
	New: func() any {
		return &models.DecodedInventory{
			Items: make([]skycrypttypes.Item, 0, 256), // Larger pre-allocation
		}
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
		return &models.DecodedInventory{
			Items: []skycrypttypes.Item{},
		}, nil
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
		buf = buf[:decodedLen]
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

	// Get pooled decoded inventory
	nbtData := decodedInventoryPool.Get().(*models.DecodedInventory)
	nbtData.Items = nbtData.Items[:0] // Reset slice but keep capacity

	decoder := nbt.NewDecoder(reader)
	_, err = decoder.Decode(nbtData)

	// Manual cleanup instead of defer (faster in hot paths)
	gzipReaderPool.Put(reader)
	bytesReaderPool.Put(bytesReader)

	if err != nil {
		decodedInventoryPool.Put(nbtData)
		return nil, err
	}

	return nbtData, nil
}
