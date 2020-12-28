package main

import (
	"fmt"
)

// This constant will shadow the other in main()
// shadowedConst in main() takes precedence over the const at package level
// This is not recommended
const shadowedConst int16 = 27

// Enumerated constants
// This are commonly handled at the package level
const enumConst = iota // evaluated to have the value 0 and inferred to have type int
// What is iota???
// iota is a counter we can use when creating enummerated constants.

// You can create a constant block here and print in main()
// iota will change the value of each const like a counter: 0, 1, 2, 3, 4...
const (
	aa = iota
	bb = iota
	cc = iota
)

// You can set aa equal to iota and leave the rest and still get the same result
// const (
// 	aa = iota
//	bb
//  cc
//)

// iota is scoped to a constant block
const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

// You can use iota with bit shifting
const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // bit shift the value by 10 * iota
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Each constant below takes up one space in a single byte
const (
	isAdmin          = 1 << iota // bit shifted 0 places so this is a literal 1
	isHeadquarters               // bit shifted 1 place so it equals 2
	canSeeFinancials             // 4

	canSeeAfrica       // 8
	canSeeAsia         // 16
	canSeeEurope       // 32
	canSeeNorthAmerica // 64
	canSeeSouthAmerica // 128
)

func main() {
	// Naming convention
	// Begin with the keyword const
	// Declare the constant the same way as a regular variable
	// If you want to export, then capitalize the first letter in the constant name

	// Typed constants
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// You CANNOT change the value of a constant
	// myConst = 27 will throw an error

	// A const has to be assignable at compile time

	// Constants can be any of the primitive types
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// Constants can be shadowed
	const shadowedConst int = 14
	fmt.Printf("%v, %T\n", shadowedConst, shadowedConst)

	// Constants work similarly to variables
	// For example, you can add a constant and variable of the same type together
	// In the case of the example below, the result will be a variable
	const c1 int = 42
	var c2 int = 27
	fmt.Printf("%v, %T\n", c1+c2, c1+c2)

	// Enumerated constants
	// This are commonly handled at the package level but you can print at function level
	// Print enumConst here
	fmt.Printf("%v, %T\n", enumConst, enumConst) // returns: 0, int

	// Print the constant block constants here
	// aa, bb, and cc are all equal to iota so they each return numbers like a counter
	// i.e. 0, 1, 2, 3,...
	fmt.Printf("%v\n", aa) // 0
	fmt.Printf("%v\n", bb) // 1
	fmt.Printf("%v\n", cc) // 2

	// from specialist constant block at package level
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist) // returns true

	// from bit shifting constant block at package level
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB) // returns 3.73GB

	// from canSee constant block at package level
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)            // returns true
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters) // returns false

	// Untyped constants

	// Enumerated expressions
}
