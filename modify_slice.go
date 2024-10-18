package main

import (
	"fmt"
)

func main() {
	sl1 := []int{1, 2, 3}
	Access_elements(sl1) // call function to access elements

	change_elements(sl1)              // call function to change elements
	fmt.Println("After change:", sl1) // print after change

	// Reassign sl1 with the appended slice
	sl1 = append_elem(sl1)
	fmt.Println("After appending elem:", sl1)
}

// Access elements of a slice
func Access_elements(n1 []int) {
	fmt.Println(n1[0]) // prints 1
	fmt.Println(n1[1]) // prints 2
}

// Change elements of a slice
func change_elements(sl []int) {
	sl[0] = 10 // change first element
	sl[1] = 20 // change second element
}

// Return the modified slice

func append_elem(sl2 []int) []int {
	// append elements and return the modified slice, slice3 = append(slice1, slice2...)
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3)
	return append(sl2, 10, 20, 30)

}
