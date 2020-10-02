package main

import (
	"fmt"

	"github.com/fbv/state"
)

const secret = "5096a665-4639-4d89-b832-6620a8a05962"

func main() {
	s := state.Generate(secret)
	fmt.Printf("State is: %s\n", s)
	fmt.Printf("Check result with wrong secret: %v\n", state.Valid(s, "wrong"))
	fmt.Printf("Check result with wrong state: %v\n", state.Valid("123", secret))
	fmt.Printf("Check result with valid state: %v\n", state.Valid(s, secret))
}
