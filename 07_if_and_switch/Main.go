package main

import (
	"fmt"
)

func main() {
	// Basic if statement
	if true {
		fmt.Println("The test is true")
	}

	// Initializer syntax
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// pop, ok := statePopulations["Florida"] is the initializer part
	// The ok after the semicolon (the boolean test) is the test part to see
	// if the code inside the code block will run or not.
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// Comparison Operators
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}

	// Other operators available: <=, >=, !=
	// When working with string types, you will use the == and != but not <, >, <=, >=

	// Using Logical Operators
	// The code below doesn't allow someone to enter a num < 1 or > 100
	number2 := 50
	guess2 := 101
	if guess2 < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess2 > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess2 < number2 {
			fmt.Println("Too low")
		}
		if guess2 > number2 {
			fmt.Println("Too high")
		}
		if guess2 == number2 {
			fmt.Println("You got it!")
		}
	}

	// Short Circuiting
	// Once Go finds an || statement that is true, it doesn't need to evaluate the rest
	// of the if statement; it already knows if the || test is going to pass.
	// The same thing happens with an && test

	// Switch Statements
	switch 3 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// Switch Statements - Falling Through
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten") // <-- This will pass
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
	// You cannot have overlapping cases
	// i.e. case 1,5, 10: and case 2,4,10: --> 10 is present twice which is an error

	// You can use an initializer (tag syntax)
	switch i := 2 + 3; i { // <-- i after the ; is a tag
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// Tagless syntax
	// With this syntax, cases can overlap
	// If they do overlap, Go will run the first case
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Switch statements in Go have break statements implied so you don't need to write them,
	// although you can if needed.
	// However, there is a keyword called "fallthrough" which means you want the cases in
	// the after where they keyword is added to run but it is not common to use
	// Keep in mind that fallthrough is logicless
	// *** Only use fallthrough if you absolutely need other cases to run! ***
	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // <-- This keyword will make case 1 and 2 print
	case j >= 20: // fallthrough is logicless so even though j = 10 it will still run
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Type Switch
	var k interface{} = 0.1 // Change the type to test
	switch k.(type) {
	case int:
		fmt.Println("k is an int")
	case float64:
		fmt.Println("k is a float 64")
	case string:
		fmt.Println("k is a string")
	default:
		fmt.Println("k is another type")
	}
}
