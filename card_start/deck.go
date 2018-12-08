package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
// A function with a receiver is like a method that belongs to an instance
// if we compare this to OOP
// we create our own type, in this case a deck type is a slice of sub type string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Coeur", "Carreaux", "Pique", "Trefle"}
	cardValues := []string{"As", "Deux", "Trois", "Quatre", "Cinq", "Six", "Sept", "Huit", "Neuf", "Dix", "Valet", "Dame", "Roi"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}
	}

	return cards
}

// receiver d of type deck, method print
// attach to the type deck a print method (like (new deck()).print())
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return two separate values from one deck, two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	// Change seed randomness for Rand package
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	r := rand.New(source)

	// randomize order of a deck with Maths package
	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // between 0 and slice - 1
		d[i], d[newPosition] = d[newPosition], d[i] // swap elements
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile(filename string) deck {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
		os.Exit(404)
	}

	return deck(strings.Split(string(file), ","))
}
