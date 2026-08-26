package main

import "fmt"

func main() {
	fmt.Println()

	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Cards:")
	remainingCards.print()

	fmt.Println("--------------------")
	fmt.Println("Deck as String:")
	fmt.Println(cards.toString())

	err := cards.saveToFile("deck.txt")
	if err != nil {
		fmt.Println("Error saving deck to file:", err)
	} else {
		fmt.Println("Deck saved to deck.txt")
	}

	fmt.Println("--------------------")
	newCards := newDeckFromFile("deck.txt")
	fmt.Println("New Deck from File:")
	newCards.print()
}
