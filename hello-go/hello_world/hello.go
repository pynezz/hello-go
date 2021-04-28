package hello

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Hello struct {
	Hello string
}

func (h Hello) Hello_Go() {
	fmt.Printf("%s, Go!", h.Hello)
}

func RandomGreeting(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	// Random format message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hello, %v. Welcome!",
		"Great to see you, %v!",
		"How you doin %v",
	}

	return formats[rand.Intn(len(formats))]
}
