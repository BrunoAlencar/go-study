package main

func main() {
	cards := newDeck()
	hand, cards := deal(cards, 5)

	hand.print()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}