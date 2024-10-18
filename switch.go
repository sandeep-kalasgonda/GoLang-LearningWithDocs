/* switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
} */

package main

import "fmt"

func main() {
	var x int = 5

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	case 5:
		fmt.Println("x is 5")

	}
}
