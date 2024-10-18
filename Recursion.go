package main

import "fmt"

func recursion(i int) int {

	if i == 11 {
		return 0
	}
	fmt.Println(i)
	return recursion(i + 1)

}

func main() {
	fmt.Println(recursion(1))

}
