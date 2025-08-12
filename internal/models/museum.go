package models

type SkyblockMuseum struct {
	Value     int64                          `json:"value,omitempty"`
	Appraisal bool                           `json:"appraisal,omitempty"`
	Items     *map[string]SkyblockMuseumItem `json:"items,omitempty"`
	Special   *[]SkyblockMuseumItem          `json:"special,omitempty"`
}

type SkyblockMuseumItem struct {
	DonatedTime  int64   `json:"donated_time,omitempty"`
	FeaturedSlot string  `json:"featured_slot,omitempty"`
	Borrowing    bool    `json:"borrowing,omitempty"`
	Items        nbtData `json:"items,omitempty"`
}
