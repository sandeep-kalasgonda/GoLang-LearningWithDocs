//nested loop

package main

import "fmt"

func main() {
	arr1 := [2]string{"sandeep", "rajendra"}
	arr2 := [3]string{"kalasgonda", "kalas", "gonda"}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			fmt.Println(arr1[i], arr2[j])
		}

	}
}
