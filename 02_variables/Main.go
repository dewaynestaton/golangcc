package main

import (
	"fmt"
	"strconv"
)

// You can declare variables at the package level here
var x string = "hello"

// You can declare multiple vars at once
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

// Shadowing is creating a var at package level and again in main()
// The var in main() takes precedence
var shadowVar string = "Testing"

func main() {
	var shadowVar string = "Testing 123"
	fmt.Println(shadowVar)

	// Three ways to declare variables in Go

	// First way...
	var i int // Use this syntax when you need to assign the variable a value later like in a loop
	i = 40
	fmt.Println(i)

	// Second way...
	var j int = 41 // Use this syntax if you need control over the type (int, bool, string, etc)
	fmt.Println(j)

	// Third way...
	k := 42
	fmt.Println(k)

	// Acronyms should be all uppercase for readability (e.g. theUrl vs theURL)
	var theURL string = "https://google.com"
	theHTTP := "https://facebook.com"
	fmt.Printf("%v\n", theURL)
	fmt.Println(theHTTP)

	// Type conversion
	var a int = 34
	fmt.Printf("%v, %T\n", a, a) // prints the value and the type

	var b float32
	b = float32(a) // this is how you convert a type
	fmt.Printf("%v, %T\n", b, b)

	// If you want to convert an int to a string, import the strconv package
	var c int = 21
	fmt.Printf("%v, %T\n", c, c)

	var d string
	d = strconv.Itoa(c)
	fmt.Printf("%v, %T\n", d, d)

	// Additional notes - Visibility
	// Lower case first letters are for package scope
	// Upper case first letters export and are available in other packages
	// No private scope
}
