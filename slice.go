package main

import (
	"fmt"
)

func main() {

	l1 := []int{}
	fmt.Println(l1)
	fmt.Println(cap(l1))
	fmt.Println(len(l1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

}

/* len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
*/
