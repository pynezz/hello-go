package basics

import (
	"fmt"
	"strings"
	"sync"
)

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

	// Anonymous struct
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

	// Iterating through all the keys and values in a map of an anonymous struct
	for k, v := range bio.friends {
		fmt.Println(k, v)
	}

	// Iterating through all the keys and values in a slice of an anonymous struct
	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}

	callFunctions()

}

// FUNCTIONS
// Functions can be values, since Go in itself is functional

func callFunctions() {
	defer LastHi() // Defer postpones a function to run last
	defer func() { // The defers works like a stack, and will be reversed in order...
		// ... adding LastHi() first, and then this one on top, executing this one before LastHi()
		fmt.Println("Almost last hi")
	}()

	a := Hello()
	fmt.Println(a) // Hello, there

	// Initializing both values short-hand
	b, c := TwoValues()
	fmt.Println(b, c) // Hello world

	// Functions as values
	d := TwoValues
	fmt.Println(d()) // Hello world

	testReceivers()
}

// Function functionName() return type { ... }
func Hello() string {
	return "Hello, there"
}

func TwoValues() (string, string) {
	return "Hello", "world"
}

// public void FunctionName()
func LastHi() {
	fmt.Println("Last HI!")
}

// RECEIVERS
// Calling a function using standard parameters will make a copy, and return the copy value.
// To mutate our values, we need to pass it along with a reference (e *Employee)
type Employee struct {
	FirstName, LastName string
}

// Standard way of sending in information to a function and specifying the return type
// function functionName(param1 type, param2 type) (returnValue returnType)
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + lastName
	return fullName
}

func testReceivers() {
	e := Employee{
		FirstName: "X Æ A-12",
		LastName:  "Musk",
	}

	fmt.Println("Son's name was", fullName(e.FirstName, e.LastName))
	e.changeFirstName("X Æ A-Xii")
	fmt.Println("Son's name is now", fullName(e.FirstName, e.LastName))
	fmt.Println("Son's name is now", e)
	testInterfaces()
}

// For quicker access to fields
func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName // More string concatination too!
}

// For modifying the internals of the struct
// I suppose this is valid:
// func (ref/ref value) funcName(p pType) (r rType, r2 rType2)
// Which makes the function a function for the struct only with params and return values
func (e *Employee) changeFirstName(firstName string) {
	e.FirstName = firstName
}

// INTERFACES
// Implementation of interfaces is implicit, rather than explicit like for C#
// Interfaces are very popular in Go

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64 // The struct implements Shape implicit
	height float64 // The struct implements Shape implicit
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width * r.height)
}

var s Shape

func testInterfaces() {
	s = Rect{width: 5.0, height: 4.0}
	r := Rect{5.0, 4.0} // We don't have to use names, just send the calues in the order of the interface

	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectangle s", s.Area())
	fmt.Println("s == r is", s == r)

	printArea(r)

	testStrings()
}

func printArea(s Shape) {
	fmt.Printf("Area of shape is: %v", s.Area())
}

// TYPE CHECKING/CONVERSION
// Everything is a type, also interfaces
func testStrings() {
	explain("Hello world")
	explain(52)
	explain(true)

	explain2("Hello world 2")

	testPointers()
}

// We can infer types at runtime
func explain(i interface{}) { // Using empty interface to infer the type of i and then checking it in the switch
	fmt.Println("Type checking with interface i as parameter: ")
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string))) // Need to cast
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func explain2(i interface{}) {
	// Shorter
	switch i := i.(type) { // Assigning i.(type) to a variable declares the type of that variable beforehand
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i)) // No need to cast
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

// POINTERS

type OriginalType struct {
	firstName string
	lastName  string
}

func changeNameVal(ot OriginalType) {
	ot.firstName = "Unchanged" // This will not change the original struct, just make a copy and return the new value
}

func changeOriginalName(ot *OriginalType) {
	ot.firstName = "Karl"
}

func changeSecondItemVal(arr [5]int) {
	arr[1] = 1
}

func changeSecondItem(arr *[5]int) {
	arr[1] = 1
}

func changeSecondItemSlice(arr []int) {
	arr[1] = 1
}

// "main"
func testPointers() {
	fmt.Println("Testing pointers")

	originalType := OriginalType{
		firstName: "Ola",
		lastName:  "Nordmann",
	}

	changeNameVal(originalType) // Do not change the original struct, only returns a copy
	fmt.Println("Did not change", originalType)

	changeOriginalName(&originalType) // Needs a receiver of the original (&) using memory address
	fmt.Println("Did change", originalType)

	// Arrays is sent by value, slices are sent by reference by default
	// Slices, maps and channels are sent by reference by default - no need to create reference
	var a = [5]int{} // Array
	changeSecondItemVal(a)
	fmt.Println("Did not change the original: ", a) // Unchanged ([0, 0, 0, 0, 0])
	changeSecondItem(&a)
	fmt.Println("Changed!", a) // This will change as it has a receiver	([0, 1, 0, 0, 0])

	var b = []int{0, 0, 0, 0}
	changeSecondItemSlice(b)
	fmt.Println("Changed clice by val..?", b) // Changed, even though we sent by value, not reference ([0, 1, 0, 0, 0])
	fmt.Println("b changed because it's a slice, and slices, maps, and channels are always sent by reference, not copy-value. Arrays sends by value")

	testGoroutines()
}

// GOROUTINES AND CHANNELS
// Goroutines are functions declared with go functionName()
// The for loops will all start a new goroutine in the background, running them separately,
// not returning 0123456789, but still correct behavior
func printHi() {
	fmt.Println("Hi")
}
func testGoroutines() {
	fmt.Println("Testing Goroutines: \"Hi\" is a return value of a goroutine 'printHi()'")
	go printHi()
	fmt.Println("Hi2")
	// Hi will sometimes not be printed...
	var wg sync.WaitGroup

	wg.Add(10)

	// This returns weird results
	// because it uses the i in the for loop as a copy
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print(i) // 101010101033101010
			wg.Done()
		}() // returnvalue
	}
	wg.Wait()

	// Testing again
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Print(i) //7126403985 all numbers from 0-9 printed
			wg.Done()
		}(i) // returnvalue
	}
	wg.Wait()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go printNum(i, &wg) // 4207186395
	}
	wg.Wait()

	//
	fmt.Println("Always pass in a parameter in the goroutine")

	testChannels()
}

func printNum(i int, wg *sync.WaitGroup) { // Passing in the waitgroup by reference
	fmt.Print(i)
	wg.Done()
}

// CHANNELS
// Allows putting data inside the channels and pull data out of those channels in a thread safe way
// Good for passing information in and out of goroutines

// Channel operator: <-
// The data flows in the direction of the arrow
func testChannels() {
	fmt.Println("Testing channels")

	simple()
}

func simple() {
	somechan := make(chan int)    // Making a new channel with the make keyword, of type int
	go func() { somechan <- 1 }() // Creating a goroutine that is going to read from the channel
	a := <-somechan               // From the outside we can create a variable that reads from the channel
	fmt.Println(a)                // And then print it

	blocking()
}

// Channels are blocking
// Channels can't hold data, so I need to get the data from the channel before the application can continue
// We need a goroutine in the background getting the value from a channel immediately after I put something in the channel
// Channels are also used for blocking and syncing the flow

func blocking() {
	someChan := make(chan bool)

	// someChan <- true 	- this will be blocked
	// <-someChan 			- as well as this
	go func() { // If we removed this function, we could not read anything in the print
		someChan <- true // because nothing is trying to write to the channel
	}() // this would cause a deadlock, resulting in an error, since the application would never finish
	fmt.Println(<-someChan) // Code will start executing from here once func blocking() is called,
	// the goroutine will be running, putting a value into the channel,
	// making us able to print the value.
	// Else, the execution would wait until the channel had a value
	buffered()
}

// Buffered channels are able to hold value
func buffered() {
	someChan := make(chan bool, 1) // Declaring a channel with one slot, meaning it can hold a single value
	someChan <- true               // No block, the channel can hold one value
	fmt.Println(<-someChan)

	brokenFor()
}

func brokenFor() {
	someChan := make(chan int)

	go func() {
		// Infinite loop in the background
		for {
			fmt.Println(<-someChan)
		}
	}() // void

	someChan <- 1 // 1
	someChan <- 2 // 2
	someChan <- 3 //
	// 3 will not be printed, because after assigning 3 to the channel, the program exits

	chanRange()
}

// Q: Why can you assign 3 values to a channel here and print all the values? Can a channel hold multiple values if they're closed?
// how do the channel know that it's being closed ahead of time? Shouldn't value 1 and 2 be discarded?
// No wait, the goroutine will assign 1 to the channel, then it will block. The for loop will read channel value 1, and get blocked
// the goroutine will assign 2, and so on. I think.
func chanRange() {
	someChan := make(chan int)
	go func() {
		someChan <- 1   // 1
		someChan <- 2   // 2
		someChan <- 3   // 3
		close(someChan) // This will work
	}()
	for val := range someChan { // This will read from the channel as it's being assigned to
		fmt.Println(val)
	}

	closingChan()
}

func closingChan() {

}
