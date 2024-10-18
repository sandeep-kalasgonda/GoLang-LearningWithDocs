/*for statement1; statement2; statement3 {
	// code to be executed for each iteration
} */

package main

import "fmt"

func main() {
	//for_loop()
	//continue_stl()
	breack_key()

}

// examples
func for_loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
}

// continue keyword
func continue_stl() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// break stm
func breack_key() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
