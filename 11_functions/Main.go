package main

import "fmt"

func main() {

	// Example 1
	sayMessage("Hello Go!")

	// Example 2
	for i := 0; i < 5; i++ {
		anotherMessage("Hello Go!", i)
	}

	// Example 3
	greeting := "Hello"
	name := "Dewayne"
	sayGreeting(greeting, name)

	// Example 4 - using pointers
	greeting2 := "Hello"
	name2 := "Zyair"
	anotherGreeting(&greeting2, &name2)
	fmt.Println(name2)

	// Example 5 - variatic parameters
	//sum{1, 2, 3, 4, 5}
}

// Example 1
func sayMessage(msg string) {
	fmt.Println(msg)
}

// Example 2
func anotherMessage(msg2 string, idx int) {
	fmt.Println(msg2)
	fmt.Println("The value of the index is", idx)
}

// Example 3
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	// You can change the value of a var and it won't affect the var in main()
	name = "Tootie"
	fmt.Println(name)
	// This prints "Tootie" but "Dewayne" will still print from main()
}

// Example 4 - using pointers
// By using pointers, you can change the value in main directly instead of working with a copy
// of the name variable
func anotherGreeting(greeting2, name2 *string) {
	fmt.Println(*greeting2, *name2)
	*name2 = "Minnie"
	fmt.Println(*name2)
}

// Example 5 - variatic parameters
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}
