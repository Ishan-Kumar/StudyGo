package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 'slice' declaration of type deck, its an abstraction over type slice
//so that we can attach some functions below as per our req
type deck []string

// reciever method to print a deck
func (d deck) print() {
	// loop declaration to iterate over slice till its range
	for index, cards := range d {
		fmt.Println(index, cards)
	}
}

// funciton to return a new deck
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites { // _ is used to tell go to ignore unused variable, here in this case we are ignoring index
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suite)
		}
	}
	return cards
}

// function to return a hand and remaining deck
// slice the list from 0 to the number pass and number to the end
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to return a string
// returning simply '[]string(d)' will throw an error because
// we are expecting to return string " ", returning slice deck instead
func (d deck) toString() string {
	return strings.Join(d, ",") // this will return a string seprated by comma
}

// function to write a deck to a file
// write in a file or return error object
// covert deck type --> string --> byte, to write in a file
func (d deck) saveDeckToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Function to read from a file
func returnDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // read file and return byte slice and err obj
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // exit with status code 1 if err
	}
	s := strings.Split(string(bs), ",") // This will return a string slice,
	return deck(s)                      //convert to deck type, this can be done coz deck is of string slice type
}

// This functions is to shuffle slice values of type deck randomly
// lines 71, 72 are done to change the seed value into a source
// so that we can have true randoms generated every time
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// now generate rnd num and swap positions
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // swap postions in slice
	}
}
