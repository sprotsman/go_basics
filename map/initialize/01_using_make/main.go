package main

import "fmt"

/*

How we can initialize a map with values.

    Using make() function

The make() function, can be used to initialize a map as shown below.

*/


func main() {
	// name map has int 'keys' and string 'values'
	var names = make(map[int]string)

	// or
	// names := make(map[int]sting)

	names[0] = "Adam"
	names[1] = "Eve"

	fmt.Println(names)

	// output ->  map[0: Adam 1: Eve]
}

/*

// Using make() to initialize a map.
m := make(map[string]bool, 0)

// Using a composite literal to initialize a map.
m := map[string]bool{}

*/