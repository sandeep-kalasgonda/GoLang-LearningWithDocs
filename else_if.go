/* Syntax
if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
} */

package main

import "fmt"

func main() {
	// Declare a variable
	a := 10
	b := 20
	c := 30
	if a > b {
		fmt.Println("a is greater than b")
	} else if c > b {
		fmt.Println("c is greater that b")
	} else {
		fmt.Println("b is greater than a and c")
	}
}
