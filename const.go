package main

import "fmt"

func main() {
	const a int = 14
	const b = "sandeep" // type not decleared
	const (
		c = 10
		d = "name"
	)

	fmt.Println(a, b, c, d)

}
