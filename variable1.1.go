//Variable Declaration Without Initial Value

package main

import (
	"fmt"
)

func main() {
	var a int
	var b string
	var c bool
	var d float64

	fmt.Println(a, ",", c, ",", b, ",", d)

	// o/p value will be set to the default value of its type:0 , false ,  , 0
}
