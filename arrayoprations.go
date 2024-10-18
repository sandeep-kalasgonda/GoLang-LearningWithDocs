package main

import (
	"fmt"
)

func main() {
	var arr1 = [...]int{1, 2, 34, 5}
	fmt.Println(arr1[0])
	arr1[0] = 20
	fmt.Println(arr1)
}
