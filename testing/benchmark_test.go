package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	networth "github.com/DuckySoLucky/SkyHelper-Networth-Go"
)

var (
	benchProfile     skycrypttypes.Profile
	benchMuseum      skycrypttypes.Museum
	benchUserProfile skycrypttypes.Member
)

func init() {
	// Load test data once for all benchmarks
	file, err := os.Open("../tests/test_profile.json")
	if err != nil {
		panic("Failed to read profile: " + err.Error())
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&benchProfile); err != nil {
		panic("Failed to parse profile")
	}

	file, err = os.Open("../tests/test_museum.json")
	if err != nil {
		panic("Failed to read museum")
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&benchMuseum); err != nil {
		panic("Failed to parse museum")
	}

	benchUserProfile = benchProfile.Members["fb3d96498a5b4d5b91b763db14b195ad"]
}

func BenchmarkNetworthCalculation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		profileNWCalc, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, *benchProfile.Banking.Balance)
		if err != nil {
			b.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
		}

		_ = profileNWCalc.GetNetworth()
	}
}

func TestNetworthPerformance(t *testing.T) {
	fmt.Println("=== Networth Performance Test ===")

	// 10 warmup runs
	fmt.Println("\nWarmup runs (10 iterations):")
	for i := 0; i < 10; i++ {
		start := time.Now()

		profileNWCalc, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, *benchProfile.Banking.Balance)
		if err != nil {
			t.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
		}

		nw := profileNWCalc.GetNetworth()
		duration := time.Since(start)

		fmt.Printf("Run %2d: %v (Networth: %.0f)\n", i+1, duration, nw.Networth)
	}

	// 100 benchmark runs
	fmt.Println("\nBenchmark runs (100 iterations):")
	var totalDuration time.Duration
	var minDuration = time.Hour
	var maxDuration time.Duration

	for i := 0; i < 100; i++ {
		start := time.Now()

		profileNWCalc, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, *benchProfile.Banking.Balance)
		if err != nil {
			t.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
		}

		nw := profileNWCalc.GetNetworth()
		duration := time.Since(start)

		totalDuration += duration
		if duration < minDuration {
			minDuration = duration
		}
		if duration > maxDuration {
			maxDuration = duration
		}

		if i < 10 || i%10 == 9 {
			fmt.Printf("Run %3d: %v (Networth: %.0f)\n", i+1, duration, nw.Networth)
		}
	}

	avgDuration := totalDuration / 100

	fmt.Println("\n=== Performance Summary ===")
	fmt.Printf("Total runs: 100\n")
	fmt.Printf("Average time: %v\n", avgDuration)
	fmt.Printf("Min time: %v\n", minDuration)
	fmt.Printf("Max time: %v\n", maxDuration)
	fmt.Printf("Total time: %v\n", totalDuration)
	fmt.Printf("Operations per second: %.2f\n", float64(time.Second)/float64(avgDuration))
}
