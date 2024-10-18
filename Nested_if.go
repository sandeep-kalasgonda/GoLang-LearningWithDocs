/* if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
} */

package main

import "fmt"

func main() {
	// if condition1 {
	a := 10

	if a > 20 {
		fmt.Println("a is greater than 20")

		if a < 20 {
			fmt.Println("a is less than 20")

		}
	} else {
		fmt.Println("a is less than or equal to 20")
	}
}
