package main

func main() {
	// Creates a new Deck of cards
	cards := newDeck()
	// hand, remainingHand := deal(cards, 5)

	// Save it to file
	cards.saveToFile("my_cards")

	// Fetch cards from the file
	nc := newDeckFromFile("my_cards")
	nc.print()

}
