package main

import "fmt"

//This program simulates a spaceship making trips to different planets using functions to calculate fuel

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Fuel left: ", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 250000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Printf("Welcome to the planet %v!\n\n---\n\n", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Printf("We do not have the available fuel to fly there.\n\n---\n\n")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	fuelRemaining := fuel
	fuelCost := calculateFuel(planet)

	fuelGauge(fuel)

	fmt.Printf("Destination: %v\nFuel needed: %v\n\n", planet, fuelCost)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	fuel := 850000

	fuel = flyToPlanet("Venus", fuel)

	fuel = flyToPlanet("Mars", fuel)

	fuel = flyToPlanet("Mercury", fuel)

	fuel = flyToPlanet("Venus", fuel)

	fuelGauge(fuel)
}
