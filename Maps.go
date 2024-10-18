/* Go Maps
Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.


****** syntax ****
map[keyType]valueType
map[keyType]valueType{key1:value1, key2:value2, ...}
map[keyType]valueType{key1: value1, key2: value2, ...
}
map[keyType]valueType{key1: value1, key2: value2, ...

var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}


using the make funcation
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)


Create an Empty Map
var a map[KeyType]ValueType


Access Map Elements
value = map_name[key]


Update and Add Map Elements
map_name[key] = value


Remove Element from Map

delete(map_name, key)

Check For Specific Elements in a Map
val, ok :=map_name[key]
}


//Create Maps Using var and :=

*/

package main

import (
	"fmt"
)

func main() {
	//Create a map using var
	var a = map[int]string{1: "sandeep", 2: "kala", 3: "sgonda"}
	b := map[string]int{"s": 1, "d": 2}
	fmt.Println(a)
	fmt.Println(b)

	map_using_make()

}

func map_using_make() {
	//Create a map using make
	var a = make(map[int]string)
	b := make(map[string]int)
	a[1] = "sandeep"
	a[2] = "kala"
	a[3] = "sgonda"
	b["s"] = 1
	b["d"] = 2
	fmt.Println(a)
	fmt.Println(b)
}
