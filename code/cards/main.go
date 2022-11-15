package main

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("mycards")
	// cards.shuffle()

	cards.print()
	cards.saveToFile("my_card")

}
