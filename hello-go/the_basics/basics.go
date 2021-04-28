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
