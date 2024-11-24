// Ex 1
/*
	Given an array of integers, return a new array with each value doubled.

	For example:
	[1, 2, 3] --> [2, 4, 6]
*/

package main
import "fmt"

// Function
func Maps(x []int) []int {

	var doubledInt int
	var doubledArray []int

	for i:= 0; i < len(x); i++ {
		doubledInt = x[i] * 2
		doubledArray = append(doubledArray, doubledInt)
	}

	return doubledArray
}


// Test
func main() {

	var xTest = []int{1, 2, 3}
	var result []int
	
	result = Maps(xTest)

	fmt.Println(result)

}