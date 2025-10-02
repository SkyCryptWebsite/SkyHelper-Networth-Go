package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	networth "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func main() {
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

	fmt.Println("Preloading prices and items...")
	networth.GetPrices(true, 69420, 3)
	networth.GetItems(true, 69420, 3)

	warmupRounds := 10
	fmt.Printf("\nRunning %d warmup rounds...\n", warmupRounds)
	for i := 0; i < warmupRounds; i++ {
		_, err := networth.CalculateFromSpecifiedInventories(inputtedInventory, models.NetworthOptions{
			IncludeItemData: true,
		})
		if err != nil {
			panic("Warmup failed: " + err.Error())
		}
	}

	benchmarkRounds := 100
	fmt.Printf("\nRunning %d benchmark rounds...\n", benchmarkRounds)

	durations := make([]time.Duration, benchmarkRounds)

	for i := 0; i < benchmarkRounds; i++ {
		start := time.Now()
		_, err := networth.CalculateFromSpecifiedInventories(inputtedInventory, models.NetworthOptions{
			IncludeItemData: true,
		})
		durations[i] = time.Since(start)

		if err != nil {
			panic(fmt.Sprintf("Benchmark round %d failed: %s", i+1, err.Error()))
		}
	}

	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	var total time.Duration
	for _, d := range durations {
		total += d
	}

	mean := total / time.Duration(benchmarkRounds)
	median := durations[benchmarkRounds/2]
	p95 := durations[int(float64(benchmarkRounds)*0.95)]
	p99 := durations[int(float64(benchmarkRounds)*0.99)]
	min := durations[0]
	max := durations[benchmarkRounds-1]

	separator := strings.Repeat("=", 50)
	fmt.Println("\n" + separator)
	fmt.Println("BENCHMARK RESULTS")
	fmt.Println(separator)
	fmt.Printf("Rounds:       %d\n", benchmarkRounds)
	fmt.Printf("Mean:         %v\n", mean)
	fmt.Printf("Median:       %v\n", median)
	fmt.Printf("Min:          %v\n", min)
	fmt.Printf("Max:          %v\n", max)
	fmt.Printf("P95:          %v\n", p95)
	fmt.Printf("P99:          %v\n", p99)
	fmt.Println(separator)

	fmt.Println("\nDistribution:")
	buckets := map[string]int{
		"< 1ms":   0,
		"1-1.5ms": 0,
		"1.5-2ms": 0,
		"2-2.5ms": 0,
		"> 2.5ms": 0,
	}

	for _, d := range durations {
		ms := float64(d.Microseconds()) / 1000.0
		if ms < 1.0 {
			buckets["< 1ms"]++
		} else if ms < 1.5 {
			buckets["1-1.5ms"]++
		} else if ms < 2.0 {
			buckets["1.5-2ms"]++
		} else if ms < 2.5 {
			buckets["2-2.5ms"]++
		} else {
			buckets["> 2.5ms"]++
		}
	}

	fmt.Printf("  < 1ms:     %3d (%.1f%%)\n", buckets["< 1ms"], float64(buckets["< 1ms"])/float64(benchmarkRounds)*100)
	fmt.Printf("  1-1.5ms:   %3d (%.1f%%)\n", buckets["1-1.5ms"], float64(buckets["1-1.5ms"])/float64(benchmarkRounds)*100)
	fmt.Printf("  1.5-2ms:   %3d (%.1f%%)\n", buckets["1.5-2ms"], float64(buckets["1.5-2ms"])/float64(benchmarkRounds)*100)
	fmt.Printf("  2-2.5ms:   %3d (%.1f%%)\n", buckets["2-2.5ms"], float64(buckets["2-2.5ms"])/float64(benchmarkRounds)*100)
	fmt.Printf("  > 2.5ms:   %3d (%.1f%%)\n", buckets["> 2.5ms"], float64(buckets["> 2.5ms"])/float64(benchmarkRounds)*100)
}
