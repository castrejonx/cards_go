package main

import (
	"fmt"
	//"math"
)

/*
Spade unicode: u2660
Heart unicode: u2665
Club unicode: u2663
Diamond unicode: u2666
*/

func main() {
	// var card string = "Ace of Spades"
	card := "A \u2660"
	fmt.Println(card)
}

func newCard() string {
	return "5 \u2666"
}