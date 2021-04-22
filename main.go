package main

import (
	"fmt"

	hello "github.com/pynezz/hello_world"
)

var age int

func main() {
	RunPackageHello_Go()
}

func PrintString() {
	age = 24
	fmt.Printf("\nMy age: %d \n", age)
	added := AddTwoInts(2, 2)
	fmt.Printf("2 + 2 = %v", added)
}

func AddTwoInts(i int, j int) int {
	add := func(i int, j int) int {
		k := i + j
		return k
	}
	return add(i, j)
}

func RunPackageHello_Go() {
	h := hello.Hello{
		Hello: "Hey",
	}
	h.Hello_Go()

	PrintString()
}
