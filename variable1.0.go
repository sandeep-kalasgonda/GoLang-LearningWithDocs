//1. syntax var varable_name type = value
//2. syntax var varable_name := value note. in this case the compiler deside the type var in compile time

package main

import (
	"fmt"
)

func main() {
	var name string = "sandeep" // sting is the datatype as decleared
	var phone = 8500000         // datatype is not declared it allocated by the compiler
	age := "23"                 // same with this also as phone type inferred

	fmt.Println(name)
	fmt.Println(phone)
	fmt.Println(age)
}
