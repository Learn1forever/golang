package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//init
func init() {
	rand.Seed(time.Now().UnixNano())
}

//
func randomFormat() string {
	formats := []string{
		"Hi,%v. Welcome!",
		"Great to see you, %v!",
		"Hail,%v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

//

//Hello returns a greetings for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
