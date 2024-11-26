package maumau

import "fmt"

// Struct für Spielkarten
type Card struct {
	Suit string
	Rank string
}

// Funktion zum Erstellen einer neuen Spielkarte.
func NewCard(suit string, rank string) Card {
	// TODO: Prüfung, ob suit und rank gültig sind.
	return Card{Suit: suit, Rank: rank}
}

// Prüft, ob zwei Karten nach den Maumau-Regeln zusammenpassen.
func (c Card) Matches(other Card) bool {
	return c.Suit == other.Suit || c.Rank == other.Rank
}

// String-Ausgabe einer Karte.
func (c Card) String() string {
	var suit string
	switch c.Suit {
	case "Herz":
		suit = "♥"
	case "Karo":
		suit = "♦"
	case "Pik":
		suit = "♠"
	case "Kreuz":
		suit = "♣"
	}
	return "{" + suit + " " + c.Rank + "}"
}

// Gibt das Bild einer Karte als Liste von Strings zurück.
func (c Card) Image() []string {
	var suit string
	switch c.Suit {
	case "Herz":
		suit = "♥"
	case "Karo":
		suit = "♦"
	case "Pik":
		suit = "♠"
	case "Kreuz":
		suit = "♣"
	}
	return []string{
		"┌─────────┐",
		// left padding
		fmt.Sprintf("│%-2s       │", c.Rank),
		"│         │",
		fmt.Sprintf("│    %s    │", suit),
		"│         │",
		fmt.Sprintf("│       %2s│", c.Rank),
		"└─────────┘",
	}
}
