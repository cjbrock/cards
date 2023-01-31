package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Clubs")

	cards.printCards()
}

func newCard() string {
	return "Five of Diamonds"
}
