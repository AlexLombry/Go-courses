package main

import (
	"fmt"
	"go-courses/external/utils"
)

func main() {
	// we are going to understand how package works
	w := utils.Walk()
	fmt.Println(w)

	a := Walk()
	fmt.Println(a)
}

func Walk() string {
	return "I walk from main"
}
