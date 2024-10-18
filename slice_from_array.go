package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	slice := arr1[0:4]
	fmt.Println(slice, arr1)
}
