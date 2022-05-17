package main

import "fmt"

/*

Using literal syntax

Maps can be initialized using map literal syntax. See example below. Using the
same code as before, modify it to initialize with map literal.

*/

func main() {
	// name map has int 'keys' and string 'values'
	var names = map[int]string {
		0: "Adam",
		1: "Eve",    // last comma must be included
	}

	fmt.Println(names)

	// output ->  map[0:Adam 1:Eve]
}


/*

// Using make() to initialize a map.
m := make(map[string]bool, 0)

// Using a composite literal to initialize a map.
m := map[string]bool{}

*/