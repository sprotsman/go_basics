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

	name, exists := names[0]
	if(exists) {
		fmt.Printf("%s exists\n", name)  // output -> "Jocko exists"
	}

	// Using the "Short Statement" If (most idiomatic)
	if name, ok := names[1]; ok {
		fmt.Printf("%s exists\n", name)
	}
	// 'name' and 'ok' no longer exist after this point, keeping your code clean.

}
