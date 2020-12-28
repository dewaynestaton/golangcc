package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a // Says b is a pointer to an integer and I want to point it to a.
	fmt.Println(a, b)
	// prints 42 and the memory location of a so b isn't holding the value of a
	// the & is the "value of" operator

	// The dereferencing operator tells us what the value at the memory location is
	a2 := 42
	var b2 *int = &a2    // the * here declares a pointer to data of that type
	fmt.Println(a2, *b2) // the * here dereferences and pulls the value at the memory addy out
	// Prints 42, 42
	// Change the value of a2 and see what happens to the value of b2
	a2 = 27
	fmt.Println(a2, *b2) // Prints 27, 27 b/c they're both pointing at the same underlying data
	// You can dereference the pointer and change the value of the data
	*b2 = 14
	fmt.Println(a2, *b2) // Prints 14, 14 because it's the same data

	// Pointer Arithmetic
	// Pointer arithmetic was intentionally left out of Go based on simplicity
	// If you absolutely need to perform pointer arithmetic, then import the "unsafe" package
	x := [3]int{1, 2, 3}              // an array of 3 values
	y := &x[0]                        // points to the index of the first value in x
	z := &x[1]                        // points to the index of the second value in x
	fmt.Printf("%v %p %p\n", x, y, z) // %p is the value of the pointer

	// Often you only want to work with the pointers and don't care where the data is stored
	// You just need the ability to point to it whereever it's at.
	var ms *myStruct        // <-- pointer to myStruct
	ms = &myStruct{foo: 42} // <-- address of myStruct with a value of 42 in it
	fmt.Println(ms)         // Prints &{42}

	// You can also use the built-in new() function
	// With the new() function you CANNOT use the object initialization syntax
	// You can only initialize an empty object
	var ms2 *myStruct
	ms2 = new(myStruct)
	fmt.Println(ms2) // Prints ${0} <-- an empty object
}

type myStruct struct {
	foo int
}
