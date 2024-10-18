package main

import (
	"fmt"
)

func main() {
	var arr1 = [5]int{}              //not initialized
	var arr2 = [5]int{1, 2}          //partially initialized
	var arr3 = [5]int{1, 2, 3, 4, 5} //fully initialized
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	inti2()
	lenght()
}

//O/P
// [0 0 0 0 0]
//[1 2 0 0 0]
//[1 2 3 4 5]

func inti2() {
	var arr4 = [5]int{1: 3, 4: 60} // Initialize Only Specific Elements

	fmt.Println(arr4)
}

//- O/P
//[0 3 0 0 60]

func lenght() { //Find the Length of an Array
	var arr5 = [...]int{1, 2, 3, 4, 5, 6, 6, 7}
	fmt.Println(len(arr5))
}

//-O/P
//8
