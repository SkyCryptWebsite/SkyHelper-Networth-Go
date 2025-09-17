package main

import (
	"fmt"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/options"
)

func main() {
	// Example usage of the public NetworthOptions
	opts := options.NetworthOptions{
		NonCosmetic:     true,
		OnlyNetworth:    false,
		IncludeItemData: true,
		SortItems:       true,
		StackItems:      false,
	}

	fmt.Printf("Created NetworthOptions with NonCosmetic: %v\n", opts.NonCosmetic)
	fmt.Printf("Created NetworthOptions with SortItems: %v\n", opts.SortItems)

	// Example of how you would use it with real data:
	/*
		calculator, err := skyhelpernetworthgo.NewProfileNetworthCalculator(userProfile, museumData, bankBalance)
		if err != nil {
			log.Fatal(err)
		}

		// Now you can use the public options!
		result := calculator.GetNetworth(opts)
		fmt.Printf("Networth: %.2f\n", result.Networth)

		// Or use without options (uses defaults)
		result2 := calculator.GetNetworth()
		fmt.Printf("Networth (default): %.2f\n", result2.Networth)

		// Or use non-cosmetic calculation
		result3 := calculator.GetNonCosmeticNetworth(opts)
		fmt.Printf("Non-cosmetic Networth: %.2f\n", result3.Networth)
	*/

	fmt.Println("NetworthOptions is now publicly available!")
}
