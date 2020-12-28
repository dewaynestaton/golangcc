package main

import (
	"fmt"
)

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ { // i is scoped to the for loop
		fmt.Println(i) // Prints 0-4
	}

	fmt.Println("\n")

	// You can change the iterator to increment by 2 or by whatever number you want
	for ii := 0; ii < 5; ii = ii + 2 {
		fmt.Println(ii) // Prints 0, 2, 4
	}

	fmt.Println("\n")

	// You can initialize more than one variable at once
	for j, jj := 0, 0; j < 5; j, jj = j+1, jj+2 {
		fmt.Println(j, jj)
		// j increments by 1, jj increments by 2
	}

	// TIP: In Go, the increment (++) operation is not an expression; it is a statement
	fmt.Println("\n")

	// The var i (or whatever var) can be left out of the for loop if initialized elswhere
	i := 0
	for ; i < 5; i++ { // you have to leave the first semicolon in
		fmt.Println(i)
	}

	fmt.Println("\n")

	// You can eleminate the increment value
	// Generates an infinite loop of 0's - uncomment to test
	// for i := 0; i < 5; {
	// 	fmt.Println(i)
	// }

	// This for loop works the same as a basic for loop and looks cleaner
	k := 0
	for k < 5 {
		fmt.Println(k) // Prints 0-4
		k++
	}

	fmt.Println("\n")

	// To run a logical test - Just like a do while loop in other languages
	// with the break statement
	x := 0
	for {
		fmt.Println(x)
		x++
		if x == 5 {
			break // <-- execution of the for loop stops and exits once we reach 4
		}
	}

	fmt.Println("\n")

	// with the continue statement
	for y := 0; y < 10; y++ {
		if y%2 == 0 {
			continue // If even, continue looping until all odd nums from 0-9 are printed
		}
		fmt.Println(y) // Prints 1, 3, 5, 7, 9
	}

	fmt.Println("\n")

	// Nested For Loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j) // Prints all of the permutations of i and j
			// only breaks out of the inner for loop
		}
	}
	fmt.Println("\n")

	// Nested For Loops with Label
	// Using a label written before the for loop and calling it in the inner loop allows you
	// to break out of both loops
Loop: // <-- The label
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 { // If you only want to print the first 3 nums and stop
				break Loop // <-- breaks out of the inner AND outer for loop
			}
		}
	}

	// How to work with collections in for loops
	s := []int{1, 2, 3} // a slice of ints containing the values 1, 2, and 3
	fmt.Println(s)

	fmt.Println("\n")

	// How do you loop through a slice? You can use a for range loop
	// This is the only syntax for using the range keyword
	r := []int{1, 2, 3}
	for k, v := range r {
		fmt.Println(k, v) // We get back the k (key: the index) and the v (values in the slice)
	}

	fmt.Println("\n")

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	// If you only want the keys, omit the v
	// If you only want the values, replace k with an underscore (_)
	// i.e. for _, v := range statePopulations {...}
}
