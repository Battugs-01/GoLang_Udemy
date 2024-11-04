package main

var deckSize int

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards = newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}
