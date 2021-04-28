package basics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (k Person) Print() {
	fmt.Println("\n", k)
}

// Capital P is syntax, not style, and needed to make the function exported (public)
func PrintPerson(name string, age int) {
	person := Person{
		name,
		age,
	}
	person.Print()
}

// SOURCE: Eli Goldberg - Golang in under an hour (2021) https://www.youtube.com/watch?v=N0fIANJkwic
// LOOPS
// REGULAR FOR LOOP
func ForLoop() {
	var sum int
	fmt.Println(sum)
	for i := 0; i < 5; i++ {
		sum++
		fmt.Println(sum)
	}
	infiniteForLoop()
}

// INFINITE FOR LOOP
func infiniteForLoop() {
	fmt.Println("Infinite for loop (with break): ")
	sum := 0
	fmt.Printf("Sum: %v", sum)
	for {
		sum++ // Goes on forever
		fmt.Printf("Sum: %v", sum)
		if sum == 10 {
			fmt.Println("Break!")
			break // Breaks at 10
		}
	}
	whileLoop()
}

// WHILE LOOP
func whileLoop() {
	// Go does not have while loops, instead you use for loop
	fmt.Println("This is a 'while' loop: n = 0")
	n := 0
	for n < 5 {
		n++
		fmt.Printf("'While' loop n = %v", n)
	}

	arrayInt()
}

// Arrays
func arrayInt() {
	var a [5]int // This will be 0, 0, 0, 0, 0
	b := [5]int{10, 20, 30, 40, 50}
	b[1] = 25
	fmt.Println("a", a)

	sliceInt()
}

// Slices are dynamically sized arrays - like C# lists
func sliceInt() {
	// Slices seems to be similar to lists in C#
	var c []int // An array without defined size is a slice
	fmt.Println("c empty: ", c)

	c = []int{10, 20, 30, 40}
	fmt.Println("c not empty: ", c)

	// Slices can also be short-hand initialized
	d := []int{10, 20, 30, 40, 50}
	fmt.Println("d: ", d)

	// Slices can also be initialized like this:
	s := make([]int, 4)
	// Append takes one slice as an argument, and the elements
	s = append(s, 60, 70)
	fmt.Println("s: ", s)

	// Manipulate the slice
	// This makes the index 2 (third element) in the slice equal to
	// the index in s where the element at the index of the length of the slice
	// (last element + 1(out of bounds)) - 1. Which is the last element.
	s[2] = s[len(s)-1]
	fmt.Println("s(3-5)", s[2:5]) // Print the 3rd to the 6th element

	// Range
	// Iterate through s
	// k is key, v is value
	// key is index
	// printing each key in each value
	for k, v := range s {
		fmt.Printf("%d is %d\n", k, v)
	}

	mapExample()
}

// MAPS
// Seems to be similar to dicts in C#
var sampleMap map[string]int

func mapExample() {
	sampleMap = map[string]int{
		"Bob":   30,
		"Kevin": 24,
	}
	// Adding to the map
	sampleMap["Another person"] = 42

	currency := map[string]string{
		"NOK": "Norwegian Krone",
		"EUR": "Euro",
		"USD": "USA Dolar",
	}

	currency["GBP"] = "Great Britain Pound"
	fmt.Println("Currency GBP added to the map: ", currency)

	// Remove from the map
	delete(currency, "GBP")
	fmt.Println("Currency GBP deleted from the map: ", currency)

	// Replacing a value in the map
	currency["EUR"] = "Norwegian Krone"
	fmt.Println("Currency with EUR value replaced with NOK: ", currency)

	// Iterating through a range in the map
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}
	// Iterating through only keys
	for key := range currency {
		fmt.Printf("%v is a currency in the map", key)
	}

	testStructs()
}

// STRUCTS
// basically like a class in C# or other languages
type anotherPerson struct {
	firstName string `json:"firstName" yaml:"firstName"` // Tagging can be added for serialization
	lastName  string
	age       int
}

type animal struct {
	name            string
	characteristics []string
}

func testStructs() {
	p1 := anotherPerson{
		firstName: "Goofy",
		lastName:  "",
		age:       0,
	}
	fmt.Println("A person struct: ", p1)

	animal1 := animal{
		name: "Lion",
		characteristics: []string{
			"Eats humans",
			"Wild animal",
			"Carnivore",
		},
	}

	// Use dot(.) to access each field in the struct
	fmt.Println("Animal name: ", animal1.name)

	// Iterate through all the values in a collection
	for _, v := range animal1.characteristics {
		fmt.Printf("\t %v\n", v) // \t is regular expression for TAB: https://www.rexegg.com/regex-quickstart.html
	}

	// Promotion ...
	type herbivore struct {
		animal   animal
		eatHuman bool
	}

	// ... A struct within a struct
	herbi := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{
				"Lacks sense",
				"Eats grass",
			},
		},
		eatHuman: false,
	}

	fmt.Println("\nThis animal:")

	fmt.Println("Eats human? ", herbi.eatHuman) // False

	// Anonymous structs
	bio := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Steven",
		friends: map[string]int{
			"Tim":   12345678,
			"Adbul": 23456789,
		},
		favDrinks: []string{
			"Pepsi Max",
			"Tea",
		},
	}

	fmt.Println("Bio: ", bio.firstName)

	for k, v := range bio.friends {
		fmt.Println(k, v)
	}

	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}

}
