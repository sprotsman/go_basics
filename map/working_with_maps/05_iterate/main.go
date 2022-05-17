package main

import "fmt"

/*

Iterating over a map entries

Using range we can iterate over a map and get keys and values both.

*/

func main() {
	var names = make(map[int]string)

	names[0] = "Adam"    // indexed insertion
	names[1] = "Eve"
	names[2] = "Cain"
	names[3] = "Abel"
	names[4] = "Enosh"

	for _, name := range names {
		fmt.Println(name)
	}

	/* output ->
		Adam
		Eve
		Cain
		Abel
		Enosh
	*/
}
