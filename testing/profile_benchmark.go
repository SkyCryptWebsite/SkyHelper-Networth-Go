package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/pprof"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	networth "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func main() {
	// Load test data
	file, err := os.Open("./tests/test_profile.json")
	if err != nil {
		panic("Failed to read profile: " + err.Error())
	}
	defer file.Close()

	var profile skycrypttypes.Profile
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		panic("Failed to parse profile")
	}

	userProfile := profile.Members["fb3d96498a5b4d5b91b763db14b195ad"]

	type SpecifiedInventory map[string]skycrypttypes.EncodedItems
	inputtedInventory := map[string]skycrypttypes.EncodedItems{
		"inventory": userProfile.Inventory.Inventory,
		"armor":     userProfile.Inventory.Armor,
		"equipment": userProfile.Inventory.Equipment,
	}

	// Preload prices and items
	fmt.Println("Preloading prices and items...")
	networth.GetPrices(true, 69420, 3)
	networth.GetItems(true, 69420, 3)

	// Warmup
	fmt.Println("Warming up...")
	for i := 0; i < 50; i++ {
		_, err := networth.CalculateFromSpecifiedInventories(inputtedInventory, models.NetworthOptions{
			IncludeItemData: true,
		})
		if err != nil {
			panic("Warmup failed: " + err.Error())
		}
	}

	// Start CPU profiling
	fmt.Println("Starting CPU profile...")
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		panic("Could not create CPU profile: " + err.Error())
	}
	defer cpuProfile.Close()

	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		panic("Could not start CPU profile: " + err.Error())
	}
	defer pprof.StopCPUProfile()

	// Run benchmark
	fmt.Println("Running profiled benchmark (1000 iterations)...")
	for i := 0; i < 1000; i++ {
		_, err := networth.CalculateFromSpecifiedInventories(inputtedInventory, models.NetworthOptions{
			IncludeItemData: true,
		})
		if err != nil {
			panic(fmt.Sprintf("Benchmark iteration %d failed: %s", i, err.Error()))
		}
	}

	fmt.Println("Profile saved to cpu.prof")
	fmt.Println("Analyze with: go tool pprof -http=:8080 cpu.prof")
}
