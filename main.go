package main

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 5)

	// hand.print()
	// remainingHand.print()

	hand.saveToFile("Hand")
	remainingHand.saveToFile("remainingHand")
}
