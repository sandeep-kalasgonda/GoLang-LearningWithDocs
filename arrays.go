// synatax var arr_name = [lenght]datatype{values}
// synatax var arr_name = [...]datatype{values} here the lenght is inferred
//using := arr_name := [lenght]datatype{values}
//using := arr_name := [...]datatype{values} here the lenght is inferred

//examples

package main

import (
	"fmt"
)

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5} // here the lenght is defined we can also use [...]
	fmt.Println(arr1)
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)
	var arr3 = [...]string{"sandeep", "k"}
	fmt.Println(arr3)

	//Access Elements of an Array

	fmt.Println(arr1[1])
	fmt.Println(arr1[0])
	fmt.Println(arr3[0])
}
