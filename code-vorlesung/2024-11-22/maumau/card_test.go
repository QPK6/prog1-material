package maumau

import "fmt"

func ExampleCard_Matches() {
	card1 := NewCard("Herz", "7")
	card2 := NewCard("Kreuz", "7")
	card3 := NewCard("Herz", "8")

	fmt.Println(card1.Matches(card2))
	fmt.Println(card2.Matches(card3))
	fmt.Println(card3.Matches(card2))

	// Output:
	// true
	// false
	// false
}

func ExampleCard_String() {
	card := NewCard("Herz", "7")
	fmt.Println(card)

	// Output:
	// {♥ 7}
}

func ExampleCard_Image() {
	card1 := NewCard("Herz", "7")
	for _, line := range card1.Image() {
		fmt.Println(line)
	}

	card2 := NewCard("Pik", "10")
	for _, line := range card2.Image() {
		fmt.Println(line)
	}

	// Output:
	// ┌───────┐
	// │7      │
	// │       │
	// │   ♥   │
	// │       │
	// │      7│
	// └───────┘
}

func ExampleDeck_Shuffle() {
	deck := New32()
	fmt.Println(deck)

	deck.Shuffle()
	fmt.Println()
	fmt.Println()
	fmt.Println(deck)

	// Output:
}
