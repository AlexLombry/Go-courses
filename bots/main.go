package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	e := englishBot{}
	s := spanishBot{}

	printGreeting(e)
	printGreeting(s)
}

func (e englishBot) getGreeting() string {
	return "Hello"
}

func (s spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
