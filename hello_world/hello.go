package hello

import "fmt"

type Hello struct {
	Hello string
}

func (h Hello) Hello_Go() {
	fmt.Printf("%s, Go!", h.Hello)
}
