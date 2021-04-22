package basics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (k Person) Print() {
	fmt.Println("\n", k)
}

func printPerson(name string, age int) {
	person := Person{
		name,
		age,
	}
	person.Print()
}
