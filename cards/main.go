package main

import "fmt"

func main() {
	card:= newCard()
	 // To run these code together,  go run cards/main.go cards/state.go
	printState()
	fmt.Println(card)		
}

func newCard() string {
	return "Five of Diamonds"
}