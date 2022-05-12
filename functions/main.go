package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
)

// Single parameter examples

func greeting(n string) {
	fmt.Printf("こんにちは (Konnichiwa), %v.\n", n)
}

func goodbye(n string) {
	fmt.Printf("じゃまたね (Ja mata ne), %v.\n", n)
}

// Multiple parameters example

// two integers
func product(a, b int) {
	fmt.Println(a * b)
}

// Variable number of parameters example (variadic)

func variadic(integers ...int) {
	for _, integer := range integers {
		fmt.Println(integer)
	}
}

// Using return values

func returnIntParam(i int) int {
	return i
}

func returnStringParam(s string) string {
	return s
}

func sum(nums ...int) int {
	sum := 0
	for n := range nums {
		sum += n
	}
	return sum
}

// use r for radius and return type of float64
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Multiple return values example

func multi(a int, b string) (int, string) {
	// return the int and the string
	return (a + 1), ("Hello " + b)
}

func nameLength(name string) (string, int) {
	return name, len(name)
}

func greeting2(n string) (string, error) {
	if len(n) == 0 {
		return "", errors.New("Empty name!")
	} else {
		// return n, nil
		return "こんにちは (Konnichiwa), " + n, nil
	}
}

// using 'ok' instead (bool)
func greeting3(n string) (string, bool) {
	if len(n) == 0 {
		return "", false
	} else {
		return "こんにちは (Konnichiwa), " + n, true
	}
}

// Pointer example

func withoutPointer(i int) {
	i += 1
}
func withPointer(i *int) {
	*i = *i + 1 // dereference with *
}

// reassign values using pointer
// val2 is int pointer type
func passBy(val1 int, val2 *int) {
	val1 = 10
	*val2 = 100 // need to dereference with *
}

// Recursion 🤮 example (prefer to avoid unless truly necessary)

func factorial(n int) int {
	// recursion needs a base case, it needs to return and stop calling itself
	if n == 1 {
		return 1
	}
	// call itself ( !5 = 5 * 4 * 3 * 2 * 1 )
	return n * factorial(n-1)
}

func fibonacciRecursive(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacciRecursive(i-1) + fibonacciRecursive(i-2)
	}
}

// Pass functions to other functions

func returnIntFunction(i int, f func(int) int) int {
	return f(i)
}

func returnStringFunction(s string, str func(string) string) string {
	return str(s)
}

// 1- use a slice of strings, 2- pass a function
func traverseNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// Anonymous functions

func anonymous() {
	// pass this anonymous function to HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("reached /")
	})

	func() {
		fmt.Println("Anonymous function!")
	}() // it requires this '()' which calls the function

	// can also assign to a variable
	f := func() {
		fmt.Println("Anonymous function!")
	}
	f()
}

// Closures: returning anonymous functions from functions
//    - retains the values or variables assigned outside of itself
//    - useful if you don't want anything else to be able to mutate the closure
//      itself

// return a func that returns an int: "func() int"
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Fibonacci using closure
func fibonacciClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func main() {
	fmt.Println("Single parameters...")
	greeting("Jocko")
	greeting("Echo")

	// to run a function at the end use 'defer'
	defer goodbye("Jocko")
	defer goodbye("Echo")

	fmt.Println()

	fmt.Println("Multiple parameters...")
	fmt.Printf("product: ")
	product(4, 3)

	fmt.Println()

	fmt.Println("variadic:")
	variadic(50, 51, 60, 61, 62)

	fmt.Println()

	fmt.Println("Using return values...")
	fmt.Printf("return int param: ")
	fmt.Println(returnIntParam(200))

	fmt.Printf("return string param: ")
	fmt.Println(returnStringParam("Echo"))

	fmt.Println()

	fmt.Printf("sum: ")
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println()

	fmt.Println("area of circle:")
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	// fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)

	fmt.Println()

	fmt.Println("Multiple return values...")
	fmt.Println(multi(10, "Jocko")) // -> 11 Hello jocko
	fmt.Printf("name length: ")
	fmt.Println(nameLength("Jocko"))

	// fmt.Println(greeting2("Jocko"))
	greet2, error := greeting2("Jocko")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(greet2)
	}

	greet3, ok := greeting3("Jocko")
	if ok {
		fmt.Println(greet3)
	}

	fmt.Println()

	fmt.Println("Pointers...")
	i := 1
	fmt.Println(i) // 1
	withoutPointer(i)
	fmt.Println(i)  // 1
	withPointer(&i) // pass address
	fmt.Println(i)  // 2

	fmt.Println("reassign - pass by value:")
	a, b := 1, 2
	passBy(a, &b)     // reference address of var b
	fmt.Println(a, b) // 1 100

	fmt.Println()

	// functions can be assigned to variables
	fmt.Println("assign function to variable:")
	hello := greeting
	hello("Jocko")

	fmt.Println()

	fmt.Println("Recursion...")
	fmt.Println("factorials:")
	fmt.Println(factorial(5))
	fmt.Println(factorial(4))
	fmt.Println("Fibonacci:")
	fmt.Println(fibonacciRecursive(10))

	fmt.Println()

	fmt.Println("Pass functions to other functions...")
	fmt.Printf("%d\n", returnIntFunction(200, returnIntParam))
	fmt.Printf("%s\n", returnStringFunction("Echo", returnStringParam))
	fmt.Println()
	fmt.Println("traverse names slice")
	// (1) pass a slice of strings and (2) call the func greeting/goodbye
	traverseNames([]string{"Jocko", "Echo", "Leif"}, greeting)
	traverseNames([]string{"Jocko", "Echo", "Leif"}, goodbye)

	fmt.Println()

	fmt.Println("Anonymous...")
	anonymous()

	fmt.Println()

	fmt.Println("Closures...")
	// assign this closure to a variable
	counter1 := counter()
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter1()) // 3

	counter2 := counter()
	fmt.Println(counter2()) // 1
	fmt.Println(counter2()) // 2
	fmt.Println(counter2()) // 3

	fmt.Println()

	fmt.Println("Fibonacci:")
	// get first 10 digits of fibonacci
	fib := fibonacciClosure()
	for i := 0; i <= 10; i++ {
		// because it is a closure we need the trailing (); this prints all zeros
		// fmt.Println(fibonacci()())
		fmt.Println(fib())
	}

	fmt.Println()
}
