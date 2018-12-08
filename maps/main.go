package main

import (
	"fmt"
)

func main() {
	// va colors map[string]string
	// colors := make(map[string]string) // another way to create a map but with no values

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// add key to existing map
	colors["white"] = "#ffffff"

	// delete a key from a map
	delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
