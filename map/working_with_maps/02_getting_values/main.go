package main

import "fmt"

/*

Getting values from a map

We can get values associated with keys from a map using index notation as shown
below.

*/

func main() {
	var names = make(map[int]string)

	names[0] = "Adam"    // indexed insertion
	names[1] = "Eve"

	fmt.Println(names[1])   // output -> Eve
}