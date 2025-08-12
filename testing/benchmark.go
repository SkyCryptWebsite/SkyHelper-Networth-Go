package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	networth "github.com/duckysolucky/skyhelper-networth-go"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

func main() {
	fmt.Println("=== Networth Calculation Benchmark ===")

	// Load test data
	file, err := os.Open("./tests/test_profile.json")
	if err != nil {
		panic("Failed to read profile: " + err.Error())
	}
	defer file.Close()

	var profile models.SkyblockProfile
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		panic("Failed to parse profile")
	}

	file, err = os.Open("./tests/test_museum.json")
	if err != nil {
		panic("Failed to read museum")
	}
	defer file.Close()

	var museum models.SkyblockMuseum
	if err := json.NewDecoder(file).Decode(&museum); err != nil {
		panic("Failed to parse museum")
	}

	userProfile := profile.Members["fb3d96498a5b4d5b91b763db14b195ad"]

	// 10 warmup runs (no detailed output)
	fmt.Println("\nPerforming 10 warmup runs...")
	for i := 0; i < 10; i++ {
		profileNWCalc, err := networth.NewProfileNetworthCalculator(&userProfile, &museum, profile.Banking.Balance)
		if err != nil {
			panic("Failed to create ProfileNetworthCalculator: " + err.Error())
		}
		_ = profileNWCalc.GetNetworth(models.NetworthOptions{})
	}

	// 100 benchmark runs with detailed timing
	fmt.Println("Starting 100 benchmark runs...")
	fmt.Println()

	var totalDuration time.Duration
	var minDuration = time.Hour
	var maxDuration time.Duration
	var networthValue float64

	for i := 0; i < 100; i++ {
		start := time.Now()

		profileNWCalc, err := networth.NewProfileNetworthCalculator(&userProfile, &museum, profile.Banking.Balance)
		if err != nil {
			panic("Failed to create ProfileNetworthCalculator: " + err.Error())
		}

		nw := profileNWCalc.GetNetworth(models.NetworthOptions{})
		duration := time.Since(start)

		if i == 0 {
			networthValue = nw.Networth
		}

		totalDuration += duration
		if duration < minDuration {
			minDuration = duration
		}
		if duration > maxDuration {
			maxDuration = duration
		}

		// Print every 10th iteration
		if i%10 == 0 || i < 10 {
			fmt.Printf("Run %3d: %v\n", i+1, duration)
		}
	}

	avgDuration := totalDuration / 100

	separator := strings.Repeat("=", 50)
	fmt.Println("\n" + separator)
	fmt.Println("BENCHMARK RESULTS")
	fmt.Println(separator)
	fmt.Printf("Networth calculated: %.0f coins\n", networthValue)
	fmt.Printf("Total runs: 100\n")
	fmt.Printf("Average time: %v\n", avgDuration)
	fmt.Printf("Minimum time: %v\n", minDuration)
	fmt.Printf("Maximum time: %v\n", maxDuration)
	fmt.Printf("Total time: %v\n", totalDuration)
	fmt.Printf("Operations per second: %.2f\n", float64(time.Second)/float64(avgDuration))
	fmt.Printf("Time per operation: %v\n", avgDuration)
}
