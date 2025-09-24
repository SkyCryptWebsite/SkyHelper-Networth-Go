package lib

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"sync"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
	"github.com/Tnze/go-mc/nbt"
)

var gzipReaderPool = sync.Pool{
	New: func() any {
		return &gzip.Reader{}
	},
}

func DecodeInventory(data string) (*models.DecodedInventory, error) {
	if data == "" {
		return &models.DecodedInventory{
			Items: []skycrypttypes.Item{},
		}, nil
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, errors.New("failed to decode base64 data: " + err.Error())
	}

	return DecodeFromBytes(decodedData)
}

func DecodeFromBytes(data []byte) (*models.DecodedInventory, error) {
	reader := gzipReaderPool.Get().(*gzip.Reader)
	defer gzipReaderPool.Put(reader)

	err := reader.Reset(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("failed to reset gzip reader: " + err.Error())
	}

	var nbtData models.DecodedInventory
	decoder := nbt.NewDecoder(reader)
	if _, err := decoder.Decode(&nbtData); err != nil {
		return nil, errors.New("failed to decode NBT data: " + err.Error())
	}

	return &nbtData, nil
}
