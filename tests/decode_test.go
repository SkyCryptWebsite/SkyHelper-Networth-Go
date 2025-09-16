package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
)

func TestDecodeInventory(t *testing.T) {
	data, err := os.ReadFile("./nbt_test_data.txt")
	if err != nil {
		t.Fatalf("failed to read nbt_test_data.txt: %v", err)
	}
	str := string(data)

	start := time.Now()
	decoded, err := lib.DecodeInventory(str)
	if err != nil {
		return
	}
	end := time.Since(start)
	fmt.Printf("time taken to decode: %v\n", end)

	// Write decoded to a JSON file
	f, err := os.Create("decoded_inventory.json")
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(decoded); err != nil {
		t.Fatalf("failed to encode JSON: %v", err)
	}
}

func TestDecodeInventory_Deathstreeks(t *testing.T) {}
