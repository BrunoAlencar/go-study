package main

import "fmt"
var number2 int = 2 // this work outside of a function
number1 := 1 // this doesnt work outside of a function 

func main() {
	// var card string = "Ace of Spades"
	card:= "Ace of Spades"

	card = "Five of Diamonds"
	fmt.Println(number1)
	fmt.Println(card)		
}