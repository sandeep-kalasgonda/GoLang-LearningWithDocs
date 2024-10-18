/* func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:


*/

package main

import (
	"fmt"
)

func parameter_example(faname string) {
	fmt.Println("hii", faname+" kalasgonda ")

}

func parameter_example1(fname string, age int) {
	fmt.Println(fname, "is", age, "years old")

}

func return_example(a int, b int) int {
	return a + b

}

func return_example1(a int, b int) (result int) {
	result = a + b
	return
}

func Multiple_Return(y int, z string) (result int, name string) {
	result = y + y
	name = z + "World"
	return
}

func skipping_return(x int, name string) (result int, text1 string) {
	result = x + x
	text1 = name + "sandeep"
	return
}

func main() {
	//parameter_example("sandeep")
	//parameter_example("rajendra")
	//parameter_example1("sandeep", 25)
	//parameter_example1("rajendra", 30)
	//fmt.Println(return_example(5, 10))
	//fmt.Println(return_example1(5, 10))
	//fmt.Println(Multiple_Return(5, "Hello"))

	/*a, b := skipping_return(5, "hello")
	fmt.Println(a, b)*/

	_, b := skipping_return(5, "hello")
	fmt.Println(b) //skipping the one or more return values

}
