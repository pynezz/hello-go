package basics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (k Person) print() {
	fmt.Println("\n", k)
}

func printPerson(name string, age int) {
	person := Person{
		name,
		age,
	}
	person.print()
}
