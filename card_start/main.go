package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()

}

/*
hand, remaining := deal(cards, 5)

hand.saveToFile("my_cards.txt")
remaining.saveToFile("deck.txt")

hand.print()

file := newDeckFromFile("my_cards.txt")

file.print()
*/

/*
// arrays in go
cards := []string{"Ace of Diamonds", newCard()} // slice of type string with specified elements
// add new element to a slice (remember array is immutable, slice is mutable)
// when you append it create a news slice, copied from the original one and add the value you want
// so basicaly it's not the original one
cards = append(cards, "Six of Spades")

// iterate over a slice (show index and the current value)
// we use the := instead of simply = because for each iteration
// go is throwing away the variable i and card, so we have to re declare our variable to iterate
for i, card := range cards {
	fmt.Println(i, card)
}
*/
