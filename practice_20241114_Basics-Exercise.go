// Exercise: Basisc
package main

import "fmt"

// Define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualTimeInOven int) int {
	return OvenTime - actualTimeInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(amountOfLayers int) int {
	return 2 * amountOfLayers
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(amountOfLayers int, actualTimeInOven int) int {
	return PreparationTime(amountOfLayers) + actualTimeInOven
}

// Test
func main() {

	// Input variables
	var actualTimeInOven int = 12
	var amountOfLayers int = 4

	// Printing results
	fmt.Println("Remaining Time:", RemainingOvenTime(actualTimeInOven))
	fmt.Println("Preparation Time:", PreparationTime(amountOfLayers))
	fmt.Println("Elapsed Time:", ElapsedTime(amountOfLayers, actualTimeInOven))
}

