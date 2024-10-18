// syntax = for index, value := range array|slice|map {
// code to be executed for each iteration }
package main

import "fmt"

func main() {
	name := [5]string{"sa", "ka", "la", "ga"}
	for ind, val := range name {
		//for _, val := range name { to skip the index val
		//for ind, _ := range name { to skip the values

		fmt.Println(ind, val)
	}

}
