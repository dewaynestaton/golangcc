package main

import (
	"fmt"
)

func main() {
	// Creation
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	// The code above works but becomes cumbersome so you should create an array
	grades := [3]int{97, 85, 93} // [3] means the array can hold up to 3 elements
	fmt.Printf("Grades: %v\n", grades)

	// Arrays are contiguous in memory meaning access to elements is very fast
	// Arrays can only hold elements of the same type

	// In the array above, we actually set the size twice ([3] and {97, 85, 93})
	// A better way is to use [...] instead of [3]
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades2)

	// You can create an array of a certain size but have its values zeroed out
	var students [4]string
	fmt.Printf("Students: %v\n", students) // returns an empty array that can hold 4 elems
	// Here's how to add elems to this array
	students[0] = "Tootie"
	fmt.Printf("Students: %v\n", students) // returns [Lisa ]
	// Continue adding students
	students[1] = "Minnie"
	students[2] = "Zy Zy"
	students[3] = "Bun Bun"
	fmt.Printf("Students: %v\n", students)

	// To get a specific elem in the array
	fmt.Printf("Student #1: %v\n", students[1]) // returns Minnie

	// To determine how big the array is, use the len() function
	fmt.Printf("Number of Students: %v\n", len(students)) // returns 4

	// Array of arrays
	// The array below stores a 3x3 identity matrix (linear algebra)
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	// Another way to do this, which might be easier to read
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	// Arrays in Go are considered as values
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5 // changes the value of index 1 to 5 (2 --> 5)
	fmt.Println(a)
	fmt.Println(b)

	// An array's fixed size has to be known at compile time which limits their usefulness
	// We can use slices
	x := []int{1, 2, 3} // this is a slice. It uses empty brackets which represents an array literal
	fmt.Println(x)      // returns [1 2 3]
	// You can get the length of a slice just like an array
	fmt.Printf("Length: %v\n", len(x)) // returns Length: 3
	// Slices have a built-in function called capacity --> cap()
	fmt.Printf("Capacity: %v\n", cap(x)) // returns Capacity: 3

	// Slices are reference types - they refer to the same underlying data
	// If you have multiple slices pointing to the same underlying data, and
	// a change is made in one of the slices, it could have
	// an impact somewhere else in your application

	// There are several ways to create a slice
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // literal way
	b1 := a1[:]                                // slice of all elements
	c1 := a1[3:]                               // slice from 4th element to end
	d1 := a1[:6]                               // slice first 6 elements
	e1 := a1[3:6]                              // slice of 4th, 5th, and 6th elements
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)

	// Slices can use the built-in make() function
	// the first param in make is the type, next is length we want, last is capacity (optional)
	z := make([]int, 3, 100)
	fmt.Println(z) // the slice is empty so it will return [0 0 0]
	fmt.Printf("Length: %v\n", len(z))
	fmt.Printf("Capacity: %v\n", cap(z))
	// You can add and remove elements in a slice using the append() function
	// append() takes two or more arguments
	z = append(z, 1, 2, 3)
	fmt.Println(z)                     // returns [0 0 0 1]
	fmt.Printf("Length: %v\n", len(z)) // returns Length: 4
	fmt.Printf("Capacity: %v\n", cap(z))

	// Common things you might do with slices
	// 1. Stack Operations
	p := []int{1, 2, 3, 4, 5}
	fmt.Println(p)
	q := p[:len(p)-1] // removes the last element (5) in p
	fmt.Println(q)

	// How do you remove an element from the middle?
	// You have to concat two slices together
	p2 := []int{1, 2, 3, 4, 5}
	q2 := append(p2[:2], p2[3:]...)
	fmt.Println(q2)
	// The first slice p2[:2] is all the elems we want to go up to
	// The second slice p2[3:] is where we concat all the elems AFTER the elem we want to remove
	// You must add the spread operator (...) to spread things out and make the append func happy =^)
}
