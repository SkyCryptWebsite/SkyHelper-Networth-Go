package lib

import (
	"bytes"
	"encoding/base64"
	"io"
	"maps"
	"slices"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/Tnze/go-mc/nbt"
	"github.com/klauspost/compress/gzip"
)

// ParseToolkit returns the tools stored in a farming or hunting toolkit as items. Slots whose tool is taken out are
// skipped: that tool is in one of the player's inventories and is counted there.
func ParseToolkit(toolkit *skycrypttypes.Toolkit) []*skycrypttypes.Item {
	if toolkit == nil || !toolkit.IsUnlocked {
		return []*skycrypttypes.Item{}
	}

	items := []*skycrypttypes.Item{}
	for _, category := range slices.Sorted(maps.Keys(toolkit.Slots)) {
		slots := toolkit.Slots[category]
		for _, index := range slices.Sorted(maps.Keys(slots)) {
			if toolkit.InUse[category][index] {
				continue
			}
			extraAttributes, err := decodeExtraAttributes(slots[index].Data)
			if err != nil || extraAttributes.Id == "" {
				continue
			}
			items = append(items, toolkitItem(extraAttributes))
		}
	}
	return items
}

// decodeExtraAttributes reads a toolkit slot: base64 of gzipped NBT holding one item's ExtraAttributes.
func decodeExtraAttributes(data string) (*skycrypttypes.ExtraAttributes, error) {
	compressed, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	raw, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var extraAttributes skycrypttypes.ExtraAttributes
	if _, err := nbt.NewDecoder(bytes.NewReader(raw)).Decode(&extraAttributes); err != nil {
		return nil, err
	}
	return &extraAttributes, nil
}

// toolkitItem builds the item a toolkit slot stands for, named after the Hypixel item.
func toolkitItem(extraAttributes *skycrypttypes.ExtraAttributes) *skycrypttypes.Item {
	name := TitleCase(extraAttributes.Id)
	if item := GetItem(extraAttributes.Id); item != nil && item.Name != "" {
		name = item.Name
	}
	count := 1
	return &skycrypttypes.Item{
		Count: &count,
		Tag: &skycrypttypes.Tag{
			ExtraAttributes: extraAttributes,
			Display:         skycrypttypes.Display{Name: name, Lore: []string{}},
		},
	}
}
