package main

// import "fmt"

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()
  cards := deck{"Ace of Diamonds", newCard()}
  cards = append(cards, "Six of Spades")

  // for i, card := range cards {
  //   fmt.Println(i, card)
  // }
  cards.shuffle()
  cards.print()

}

func newCard() string {
  return "Five of Diamonds"
}
