package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.makeHandAndPrint()
	//cards.createfileWithDeck()
	//fetchDeckFromFile()
	cards.shufNewDeck()
}

// we do it like reciever function so that we can do class like call when it makes sense
func (cards deck) makeHandAndPrint() {
	cards = newDeck() // Card values will remain the same even after slicing
	hand, remaininDeck := deal(cards, 5)
	fmt.Println(" **** Hand Returned ****")
	hand.print()

	fmt.Println()
	fmt.Println(" **** Remaining Deck ****")
	remaininDeck.print()
}

func (cards deck) createfileWithDeck() {
	cards = newDeck()
	cards.saveDeckToFile("My_Cards")
}

func fetchDeckFromFile() {
	returnDeckFromFile("My_Cards").print()
}

func (cards deck) shufNewDeck() {
	cards.shuffle()
	cards.print()
}

//to run  go run (ls *.go)
