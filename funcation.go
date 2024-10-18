/*func FunctionName() {
  // code to be executed
}

Naming Rules for Go Functions
A function name must start with a letter
A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
Function names are case-sensitive
A function name cannot contain spaces
If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used
Tip: Give the function a name that reflects what the function does!

*/

package main

import "fmt"

func main() {
	function_example()
	function_example()

}

func function_example() {
	fmt.Println("Hello, World!")

}
