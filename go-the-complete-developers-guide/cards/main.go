package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := newDeckFromFile("my_cards")

	cards.print()
}
