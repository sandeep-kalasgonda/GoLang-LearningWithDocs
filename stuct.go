/* Go Structures
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records. For example, you might have a struct to represent a person, with fields for name, age,

****** syntax *********
type struct_name struct {
member1 data_type
member2 data_type
member3 data_type
...
}

*/

package main

type student struct {
	name  string
	age   int
	grade float32
	class int
}
