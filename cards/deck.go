package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// the receiver is a reference to the actual copy of the deck we're working with
// by convention we use a one or two letter abbreviation of the type
func (d deck) print () {
	for i, card := range d {
		// if any variable is not used go will throw an error
		fmt.Println(i, card)
	}
}