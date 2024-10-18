//Create a Slice With The make() Function

package main

import "fmt"

func main() {
	// Create a slice with the make() function  slice_name := make([]type, length, capacity)
	sl1 := make([]int, 5, 10)
	sl2 := make([]int, 5) // without cap
	fmt.Println(sl1, len(sl1), cap(sl1))
	fmt.Println(sl2, len(sl2), cap(sl2))
}

//O/P
// [0 0 0 0 0] 5 10
// [0 0 0 0 0] 5 5
