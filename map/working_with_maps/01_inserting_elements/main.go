package main

import "fmt"

/*

Inserting elements in a map

You can insert keys in two ways. Either insert keys when initializing or use
index syntax to initialize.

*/

func main() {
	var names = make(map[int]string)

	names[0] = "Adam"    // indexed insertion
	names[1] = "Eve"

	fmt.Println(names)   // output -> map[0:Adam 1:Eve]
}