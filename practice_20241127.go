// Ex 1
/*
	Return a new array consisting of elements which are multiple of their own index in input array (length > 1).
	
	Some cases:
		[22, -6, 32, 82, 9, 25] =>  [-6, 32, 25]
		[68, -1, 1, -7, 10, 10] => [-1, 10]
		[-56,-85,72,-26,-14,76,-27,72,35,-21,-67,87,0,21,59,27,-92,68] => [-85, 72, 0, 68]
*/

package main
import "fmt"

// Function
func elementsMultiple(elements []int) []int {
	// Output object
	var select_element []int

	for i := 1; i < len(elements); i++ {
		if elements[i] % i == 0 {
			select_element = append(select_element, elements[i])
		}
	}
	return select_element
}

// Test
func main() {

	var test_array = []int {22, -6, 32, 82, 9, 25}
	var result []int

	result = elementsMultiple(test_array)

	fmt.Println(result)
}
