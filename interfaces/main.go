package main

import "fmt"

// Interfaces are created to reuse functions
// Notice printGreeting function can be used to print both eb, sb struct types
// Because getGreeting functinos are types eb & eb hence both structs became the part of interface

type bot interface {
	getGreeting() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printGeeting(eb)
	printGeeting(sb)
}

// this print function can now be reused to print any type bot
func printGeeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishbot) getGreeting() string {
	// custom logic for english
	return "hi there!"
}

func (spanishbot) getGreeting() string {
	// custom logic for spanish
	return "hola!"
}
