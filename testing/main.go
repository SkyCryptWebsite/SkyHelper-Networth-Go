package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	networth "duckysolucky/skyhelper-networth-go"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

func main() {
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

	profileNWCalc, err := networth.NewProfileNetworthCalculator(&userProfile, &museum, profile.Banking.Balance)
	if err != nil {
		panic("Failed to create ProfileNetworthCalculator: " + err.Error())
	}

	timeNow := time.Now()
	nw := profileNWCalc.GetNonCosmeticNetworth()
	fmt.Printf("Time: %s\n", time.Since(timeNow))
	fmt.Printf("Networth: %+v\n", nw.Types)

	f, err := os.Create("networth.json")
	if err != nil {
		panic("Failed to create file: " + err.Error())
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(nw); err != nil {
		panic("Failed to encode JSON: " + err.Error())
	}

	fmt.Printf("Networth: %+v\n", nw.UnsoulboundNetworth)
}
