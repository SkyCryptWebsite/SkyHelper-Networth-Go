package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	networth "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/options"
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

	file, err = os.Open("./tests/test_museum.json")
	if err != nil {
		panic("Failed to read museum")
	}
	defer file.Close()

	var museum skycrypttypes.Museum
	if err := json.NewDecoder(file).Decode(&museum); err != nil {
		panic("Failed to parse museum")
	}

	userProfile := profile.Members["fb3d96498a5b4d5b91b763db14b195ad"]

	profileNWCalc, err := networth.NewProfileNetworthCalculator(&userProfile, &museum, *profile.Banking.Balance)
	if err != nil {
		panic("Failed to create ProfileNetworthCalculator: " + err.Error())
	}

	timeNow := time.Now()
	nw := profileNWCalc.GetNonCosmeticNetworth(options.NetworthOptions{IncludeItemData: false})
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

	fmt.Printf("Networth: %+v\n", nw.Types["inventory"].Items)

	/*
	prices, err := networth.GetPrices(true, 69420, 3)

	count := 1

	calculatorService := networth.NewCalculatorService()
	itemNWCalc := calculatorService.NewSkyBlockItemCalculator(
		&skycrypttypes.Item{
			Tag: &skycrypttypes.Tag{
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Id: "HYPERION",
				},
			},
			Count: &count,
		},
		*prices,
		options.NetworthOptions{},
	)

	calculatorService.CalculateItem(itemNWCalc)

	fmt.Printf("\n\nPrice: %+v\n", itemNWCalc.Price)
	*/

}
