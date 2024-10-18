/*
switch expression {
case x,y:
   // code block if expression is evaluated to x or y
case v,w:
   // code block if expression is evaluated to v or w
case z:
...
default:
   // code block if expression is not found in any cases
}
*/

package main

import "fmt"

func main() {
	var x int = 10

	switch x {
	case 5, 3:
		fmt.Println("x is 5 or 3")
	case 10, 1, 4:
		fmt.Println("x is 10, 1 or 4")
	default:
		fmt.Println("x is not 5, 3, 10, 1 or")
	}
}
