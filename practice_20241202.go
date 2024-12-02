// Ex 1
/*
	I have a cat and a dog.
	I got them at the same time as kitten/puppy. That was humanYears years ago.
	Return their respective ages now as [humanYears,catYears,dogYears]

	NOTES:
		humanYears >= 1
		humanYears are whole numbers only

	Cat Years
		15 cat years for first year
		+9 cat years for second year
		+4 cat years for each year after that

	Dog Years
		15 dog years for first year
		+9 dog years for second year
		+5 dog years for each year after that
*/

package main
import "fmt"

// Function
func humanCatDogYears(years int) (result [3]int) {
	var cat_age, dog_age int

	if years <= 1 {
		cat_age = 15
		dog_age = 15
	} else if years <= 2 {
		cat_age = 15 + 9
		dog_age = 15 + 9
	} else {
		cat_age = 24 + (years - 2) * 4
		dog_age = 24 + (years - 2) * 5
	}

	result = [3]int {years, cat_age, dog_age}

	return result
}


// Test

