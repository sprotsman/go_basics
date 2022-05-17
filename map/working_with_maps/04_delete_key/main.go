package main

import "fmt"

/*

Delete key from a map

We use the delete function to remove a key from a map. Let’s see an example of
how it is done.

*/

func main() {
	var names = make(map[int]string)

	names[0] = "Jocko"    // indexed insertion
	names[1] = "Echo"
	names[2] = "Aurelius"

	// delete Cain
	delete(names, 2)
	fmt.Println(names)   // prints map[0:Jocko 1:Echo]
}
