package models

import skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"

type DecodedNewYearCakeBagData struct {
	Items []skycrypttypes.Item `nbt:"i" json:"items,omitempty"`
}

type DecodedDisplay struct {
	Name  string   `nbt:"Name" json:"name,omitempty"`
	Lore  []string `nbt:"Lore" json:"lore,omitempty"`
	Color int      `nbt:"color" json:"color,omitempty"`
}

type DecodedInventory struct {
	Items []skycrypttypes.Item `nbt:"i" json:"items,omitempty"`
}
