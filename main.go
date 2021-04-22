package main

import (
	"fmt"

	hello "github.com/pynezz/hello_world"
	basics "github.com/pynezz/the_basics"
)

var age int
var name string = "Kevin"

func main() {
	// runPackageHello_Go()
	// printFromBasics(name, 24)
	scraper()
}

func printFromBasics(name string, age int) {
	basics.PrintPerson(name, age)
}

func printString() {
	age = 24
	fmt.Printf("\nMy age: %d \n", age)
	added := addTwoInts(2, 2)
	fmt.Printf("2 + 2 = %v", added)
}

func addTwoInts(i int, j int) int {
	add := func(i int, j int) int {
		k := i + j
		return k
	}
	return add(i, j)
}

func runPackageHello_Go() {
	h := hello.Hello{
		Hello: "Hey",
	}
	h.Hello_Go()

	printString()
}

// Notes:
