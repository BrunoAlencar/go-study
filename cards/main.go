package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// hand.saveToFile("cards/hand.txt")

	cards2 := newDeckFromFile("cards/hand.txt")
	cards2.print()

}
