package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

// func main() {
// 	// var card string = "Ace of Spades"
// 	card := newCard()

// 	// var keyword, plus the name, and then the type.
// 	// only a value of type string will be assigned to the value. (statically typed language)
// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }
