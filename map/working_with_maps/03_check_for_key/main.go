package main

import "fmt"

/*

Check if a key exists

We get a boolean when we try retrieving a value. That helps to check if a key
exists.

*/

func main() {
	var names = make(map[int]string)

	names[0] = "Jocko"    // indexed insertion
	names[1] = "Echo"

	adam, exists := names[0]
	if(exists) {
		fmt.Printf("%s exists\n", adam)  // output -> "Adam exists"
	}
}