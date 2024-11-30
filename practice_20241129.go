// Ex 1
/*
	Create a method each_cons that accepts a list and a number n, 
	and returns cascading subsets of the list of size n, like so:

	each_cons([1,2,3,4], 2)
	#=> [[1,2], [2,3], [3,4]]

	each_cons([1,2,3,4], 3)
	#=> [[1,2,3],[2,3,4]]
	As you can see, the lists are cascading; ie, they overlap, but never out of order.
*/

package main
import "fmt"

// Function
func eachCons(arr []int, n int) [][]int {
	// Support array
	var out_array [][]int

	for i:= 0; i <= len(arr)-n; i++ {
		var out_array_element []int
		for j := 0; j < n; j++ {
			out_array_element = append(out_array_element, arr[i+j])
		}
		out_array = append(out_array, out_array_element)
	}
	return out_array
}

// Test
func main() {
	var test_array = []int {1, 2, 3, 4}
	var result = eachCons(test_array, 2)

	fmt.Println(result)
}
