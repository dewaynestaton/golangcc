package main

import "fmt"

func main() {
	// Boolean type
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1
	o := 1 == 2
	fmt.Printf("%v, %T\n", m, m) // true
	fmt.Printf("%v, %T\n", o, o) // false

	// bool's zero value is false
	var p bool // automatically set to false unless you make it equal true
	fmt.Printf("%v, %T\n", p, p)

	// Numeric types
	// (1) Integers
	// The zero value for ALL numerical types is 0
	// An integer is going to be AT LEAST 32 bits (default)

	// Below are signed integers...
	// int8  --> -128 - 127
	// int16 --> -32,768 - 32,767
	// int32 --> -2,147,483,648 - 2,147,483,647 (billions)
	// int64 --> -9,233,372,036,854,755,808 - 9,223,372,036,854,775,807 (quintillions)
	// Any number bigger than the above is for a huge program and will require using a package

	// Below are unsigned integers
	// uint8  --> 0 - 255 // very common type
	// uint16 --> 0 - 65,535
	// uint32 --> 0 - 4,294,967,295
	// Unsigned integers are always positive!

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // Interger division returns an int and drops the remainder. This returns 3.
	fmt.Println(a % b) // This will return the remainder of 1

	// You CANNOT add integers of different types
	// Uncomment and see that this returns an error
	//var y int = 10
	//var z int8 = 3
	//fmt.Println(y + z)

	// You CAN do a type conversion to get it to work
	var y int = 10
	var z int8 = 3
	fmt.Println(y + int(z))

	// Another operation uses bit operators
	s := 10             // 1010 <-- binary representation of 10
	t := 3              // 0011
	fmt.Println(s & t)  // AND --> 10 AND 3 --> 0010 == 2
	fmt.Println(a | t)  // OR --> 10 OR 3 --> 1011 == 11
	fmt.Println(s ^ t)  // EXCLUSIVE OR --> 10 OR 3 has the bit set but not both --> 1001 == 9
	fmt.Println(s &^ t) // AND NOT --> Opposite of ^. Set true only if neither ints have bit set  --> 0100 == 8

	// Another operation is called bit shifting
	h := 8              // Same as 2^3
	fmt.Println(h << 3) // Shifts h LEFT three places - same as 2^3 * 2^3 = 2^6 = 64
	fmt.Println(h >> 3) // Shifts h RIGHT three places - same as 2^3 / 2^3 = 2^0 = 1

	// (2) Floating point
	// Follows IEEE 754 Standard
	// float32 +-1.18E38 --> +-3.4E38 (meaning + or - 1.18 * 10^38)
	// If you need more precision than float32 use
	// float64 +-2.23E308 --> +-1.8E308
	e := 3.14
	e = 13.7e72 // 13.7 * 10^72
	e = 2.1e14  // 2.1 * 10^14
	fmt.Printf("%v, %T\n", e, e)

	f := 10.2
	g := 3.7
	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)

	// (3) Complex numbers
	// Two types of complex numbers: complex64 and complex128
	var q complex64 = 1 + 2i // Go recognizes i as an imaginary number
	fmt.Printf("%v, %T\n", q, q)

	r := 1 + 2i
	rr := 2 + 5.2i
	fmt.Println(r + rr)
	fmt.Println(r - rr)
	fmt.Println(r * rr)
	fmt.Println(r / rr)

	// What happens if you need to decompose this down and get
	// the real part or the imaginary part?
	// You can use the real() func or the imag() func
	var qq complex64 = 1 + 5.2i
	fmt.Printf("%v, %T\n", real(qq), real(qq))
	fmt.Printf("%v, %T\n", imag(qq), imag(qq))
	// If you run real() or imag() on complex64's, you'll get float32's
	// If you run the functions on complex128's, you'll get float64's

	// If you need to make a complex number, you can use the complex() func
	var pp complex128 = complex(5, 12) // will create (5 + 12i) for you
	fmt.Printf("%x, %T\n", pp, pp)

	// Text types
	// The first text type is a string and stands for any UTF-8 character
	// Strings CANNOT encode every UTF-8 char and will require a rune

	word := "this is a string"
	fmt.Printf("%v, %T\n", word, word)

	// Strings can be looked at like a collection of letters (think of an array)
	// so you can pull out the index of any character
	word2 := "this is a string"
	fmt.Printf("%v, %T\n", string(word2[2]), word2[2]) // gets the third index which is the letter i

	// You can concat like this
	w3 := "This is a string."
	w4 := " This is also a string."
	fmt.Printf("%v, %T\n", w3+w4, w3+w4) // reminder that %T just tells you the data type

	// In Go, you can create a collection of bytes aka slices of bytes
	w3 = "Yup, this is a string"
	w3Byte := []byte(w3)
	fmt.Printf("%v, %T\n", w3Byte, w3Byte) // returns the ASCII values

	// Rune - represents any UTF-32 character
	// When using a rune, use single quotes
	// A rune is just a type alias for int32's
	r1 := 'a'
	fmt.Printf("%v, %T\n", r1, r1) // returns 97

	var r2 rune = 'a'
	fmt.Printf("%v, %T\n", r2, r2) // returns the same results as above
}
