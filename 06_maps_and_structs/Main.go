package main

import (
	"fmt"
	"reflect"
)

// Struct example - go to line 74
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// the keyword 'struct' is added as the type but it can be written differently
// type Doctor struct {} which is the regular way or...
// use an anonymous struct
// aDoctor := struct{name string}{name: "John Pertwee"}
// Then you can print aDoctor to get the same result
// This is a short-lived struct (used maybe only one time)

// Embedding example - go to line 112
type Animal struct {
	Name   string `required max: "100"` // <-- This is a tag which describes the field
	Origin string
}

type Bird struct {
	Animal   // embeds the Animal struct
	SpeedKPH float32
	CanFly   bool
}

type Cat struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// MAPS
	// maps a key (the string) to a value (the int)
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations, "\n")

	// With maps, an array is a value key type, but a slice is not

	// The slice below will throw an error
	//m := map[[]int]string{}
	//fmt.Println(statePopulations, m)

	// The array below will work
	m2 := map[[3]int]string{}
	fmt.Println(statePopulations, m2, "\n") // returns an empty map along with population map

	// There are other ways to create a map besides the literal syntax
	// We can use the make() func
	statePopulations2 := make(map[string]int)
	statePopulations2 = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations2, "\n")

	// To pull out a piece of data like the population of Cali
	fmt.Println(statePopulations2["California"]) // returns 39250017
	// Adding a key/value pair to a map
	statePopulations2["Georgia"] = 10310371
	fmt.Println(statePopulations2, "\n")

	// Tip: In a map, things are not always returned in the order you might expect

	// You are able to use the built-in delete() func
	delete(statePopulations2, "Georgia")
	fmt.Println(statePopulations2, "\n")

	// If you're not sure if the key/value is in the map or not, use the ok keyword
	pop, ok := statePopulations["Oregon"]
	fmt.Println(pop, ok) // returns 0, false because the key was not found
	// If you just want a boolean value to tell if a key/val exists
	_, ok = statePopulations["Ohio"]
	fmt.Println(ok) // returns true because Ohio is in the map

	// You can find out how many elements are in a map using the len() func
	fmt.Println(len(statePopulations))

	/////////////////////////////////////////////////////////////////////////////////////////////

	// STRUCTS
	// Structs can have different data types and other structs
	// When naming, you can use upper case for the keys so they can be exported
	// (i.e. Number: 3, ActorName: "Jon Pertwee")
	// Just remember they are capitalized when creating the struct outside of main()
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	// To get one item from the struct like the actorName
	fmt.Println(aDoctor.actorName) // returns Jon Pertwee
	// You can get the slice or a slice of the slice
	// The slice
	fmt.Println(aDoctor.companions) // returns the entire slice
	// The slice of a slice
	fmt.Println(aDoctor.companions[1]) // returns Jo Grant

	// Pointing to a struct
	anotherDoctor := &aDoctor // anotherDoctor is a pointer to aDoctor
	anotherDoctor.actorName = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	// Embedding
	// Go does not have an inheritance model but it supports
	// composition through embedding (has a - Bird has an Animal)
	// In Java, this would be Bird is an Animal (inheritance)
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	// If using the literal syntax - works the same as the Bird example
	c := Cat{
		Animal:   Animal{Name: "Serval", Origin: "Africa"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	fmt.Println(c)

	// Tags (see the tag in the Animal struct)
	// Requires you to import the "reflect" package
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // returns "required max: 100"
}
