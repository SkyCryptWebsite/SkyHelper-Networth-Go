package main

import (
	"encoding/json"
	"os"
	"testing"

	networth "duckysolucky/skyhelper-networth-go"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

var (
	benchProfile     models.SkyblockProfile
	benchMuseum      models.SkyblockMuseum
	benchUserProfile models.SkyblockProfileMember
)

func init() {
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

	for b.Loop() {
		profileNWCalc, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, benchProfile.Banking.Balance)
		if err != nil {
			b.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
		}

		_ = profileNWCalc.GetNetworth(models.NetworthOptions{})
	}
}

func BenchmarkNetworthCalculatorCreation(b *testing.B) {

	for b.Loop() {
		_, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, benchProfile.Banking.Balance)
		if err != nil {
			b.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
		}
	}
}

func BenchmarkNetworthCalculationOnly(b *testing.B) {
	profileNWCalc, err := networth.NewProfileNetworthCalculator(&benchUserProfile, &benchMuseum, benchProfile.Banking.Balance)
	if err != nil {
		b.Fatalf("Failed to create ProfileNetworthCalculator: %v", err)
	}

	for b.Loop() {
		_ = profileNWCalc.GetNetworth(models.NetworthOptions{})
	}
}
