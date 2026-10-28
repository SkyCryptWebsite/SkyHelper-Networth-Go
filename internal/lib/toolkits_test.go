package lib

import (
	"bytes"
	"encoding/base64"
	"testing"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/Tnze/go-mc/nbt"
	"github.com/klauspost/compress/gzip"
)

func encodeToolkitSlot(t *testing.T, id string) skycrypttypes.EncodedItems {
	t.Helper()
	data, err := nbt.Marshal(map[string]any{"id": id})
	if err != nil {
		t.Fatal(err)
	}
	var compressed bytes.Buffer
	writer := gzip.NewWriter(&compressed)
	_, _ = writer.Write(data)
	_ = writer.Close()
	return skycrypttypes.EncodedItems{Data: base64.StdEncoding.EncodeToString(compressed.Bytes())}
}

// Ported from SkyHelper-networth's parseToolkit tests.
func TestParseToolkit(t *testing.T) {
	toolkit := &skycrypttypes.Toolkit{
		IsUnlocked: true,
		InUse:      map[string]map[int]bool{"TRAP": {0: true}},
		Slots: map[string]map[int]skycrypttypes.EncodedItems{
			"TRAP":           {0: encodeToolkitSlot(t, "TRAP_IN_USE"), 1: encodeToolkitSlot(t, "TRAP")},
			"HUNTING_SCYTHE": {0: encodeToolkitSlot(t, "HUNTING_SCYTHE")},
			"LASSO":          {0: {Data: "not base64 nbt"}},
		},
	}
	items := ParseToolkit(toolkit)
	ids := map[string]bool{}
	for _, item := range items {
		if item.Count == nil || *item.Count != 1 || item.Tag == nil || item.Tag.Display.Name == "" {
			t.Fatalf("toolkit item not filled in: %+v", item)
		}
		ids[item.Tag.ExtraAttributes.Id] = true
	}
	if len(items) != 2 || !ids["TRAP"] || !ids["HUNTING_SCYTHE"] {
		t.Fatalf("got %v, want TRAP and HUNTING_SCYTHE (the in-use trap and the broken slot skipped)", ids)
	}

	if got := ParseToolkit(&skycrypttypes.Toolkit{IsUnlocked: false, Slots: toolkit.Slots}); len(got) != 0 {
		t.Fatalf("a locked toolkit gave %d items", len(got))
	}
	if got := ParseToolkit(nil); got == nil || len(got) != 0 {
		t.Fatal("a missing toolkit should give an empty list")
	}
}
