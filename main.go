package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("-----")

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// fmt.Println("-----")

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println("-----")

	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// fmt.Println("-----")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
