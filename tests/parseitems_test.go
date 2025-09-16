package tests

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
)

func TestParseItems(t *testing.T) {
	file, err := os.Open("./test_profile.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var profile skycrypttypes.Profile
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		t.Fatal(err)
	}

	file, err = os.Open("./test_museum.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var museum skycrypttypes.Museum
	if err := json.NewDecoder(file).Decode(&museum); err != nil {
		t.Fatal(err)
	}

	start := time.Now()
	member := profile.Members["fb3d96498a5b4d5b91b763db14b195ad"]
	decoded, err := lib.ParseItems(&member, &museum)
	duration := time.Since(start)
	t.Logf("ParseItems took %s", duration)

	if err != nil {
		t.Error("Expected ParseItems to succeed but got error:", err)
	}

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

func BenchmarkParseItems(b *testing.B) {
	file, err := os.Open("./test_profile.json")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	var profile skycrypttypes.Profile
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		member := profile.Members["fb3d96498a5b4d5b91b763db14b195ad"]
		_, err := lib.ParseItems(&member, nil)
		if err != nil {
			b.Error("Expected ParseItems to succeed")
		}
	}
}
