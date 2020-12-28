package main

import (
	"fmt"
	"log"
)

func main() {
	// Defer
	// fmt.Println("start")
	// fmt.Println("middle")
	// fmt.Println("end")

	// fmt.Println("\n")

	// Defer will make a function or statement execute at the last
	//fmt.Println("start")
	//defer fmt.Println("middle") // <-- this will print last
	//fmt.Println("end")

	//fmt.Println("\n")

	// Defer statements are in LIFO order
	// These will print in the reverse order
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Printf("%s\n", robots)

	// TIP: If working with a for loop, it may not be the best idea to defer a Close() function.

	// Panic
	// a, b := 1, 0
	// ans := a / b     // 1 / 0 is invalid
	// fmt.Println(ans) // The terminal will create a panic message for you in this case

	// Here's how to use the built-in panic function
	// fmt.Println("start")            // this prints
	// panic("something bad happened") // a panic message will print
	// fmt.Println("end")              // this will not print

	// A more practical example
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// 	// if you run this in the terminal more than once it will throw a panic error message
	// }

	// Panics happen AFTER deferred statements are executed
	// fmt.Println("start")                   // prints normally
	// defer fmt.Println("this was deferred") // this prints 2nd even though it was deferred
	// panic("something bad happened")        // this prints last
	// fmt.Println("end")                     // this never is reached

	// Recover
	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // <-- the recover runs after the panic
			log.Println("Error:", err)
		}
	}() // <-- this is what will make the anon function run
	panic("something bad happened")
	fmt.Println("done panicking") // never reached because the program panicked
}
