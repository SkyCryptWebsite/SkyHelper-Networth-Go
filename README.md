# SkyHelper-Networth-Go

[![discord](https://img.shields.io/discord/720018827433345138?logo=discord)](https://discord.com/invite/fd4Be4W)
[![license](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/github.com/DuckySoLucky/SkyHelper-Networth-Go.svg)](https://pkg.go.dev/github.com/github.com/DuckySoLucky/SkyHelper-Networth-Go)
> [!NOTE] 
> This repository is specifically designed for integration with [SkyCryptv3](https://github.com/DuckySoLucky/SkyCryptv3/). While this Go implementation focuses on core networth calculation functionality (including standard and non-cosmetic calculations), it does not include all features available in the Node.js module, such as UpdateManager and NetworthManager components.

[SkyHelper](https://skyhelper.altpapier.dev/)'s Networth Calculation as a Go module to calculate a player's SkyBlock networth by using their profile data provided by the [Hypixel API](https://api.hypixel.net/).

## Installation

```bash
go get github.com/DuckySoLucky/SkyHelper-Networth-Go
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/DuckySoLucky/SkyHelper-Networth-Go"
)

func main() {
    userProfile := // https://api.hypixel.net/#tag/SkyBlock/paths/~1v2~1skyblock~1profile/get - profile.Members[uuid]
    museumData := // https://api.hypixel.net/v2/skyblock/museum - museum.Members[uuid]

	calculator, err := skyhelpernetworthgo.NewProfileNetworthCalculator(userProfile, museumData, profile.Banking.Balance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create networth calculator: %v", err),
		})
	}

    networth := calculator.GetNetworth()
    nonCosmeticNetworth := calculator.GetNonCosmeticNetworth()

    fmt.Printf("Networth: %+v\n", networth)
    fmt.Printf("Non-Cosmetic Networth: %+v\n", networth)
}
```


## Testing

Run the comprehensive test suite:

```bash
go test ./tests

# Run specific item tests
go test ./tests/ -run TestEnchantedBook
go test ./tests/ -run TestPets
go test ./tests/ -run TestGemstones

# Run with verbose output
go test -v ./tests/...
```

## Benchmarking

Performance benchmarks are available to measure calculation speed:

```bash
# Run benchmarks
go test -bench=. ./benchmark/...

# Run specific benchmarks
go test -bench=BenchmarkNetworth ./benchmark/...

# Run benchmarks with memory profiling
go test -bench=. -benchmem ./benchmark/...
```

## License

This project is licensed under the terms specified in [LICENSE.MD](LICENSE.MD).
