package maumau

import "math/rand"

type Deck []Card

// New32 liefert ein neues Skat-Blatt mit 32 Karten.
func New32() Deck {
	deck := make(Deck, 32)
	for s, suit := range []string{"Karo", "Herz", "Pik", "Kreuz"} {
		for r, rank := range []string{"7", "8", "9", "10", "Bube", "Dame", "König", "Ass"} {
			card := Card{suit, rank}
			deck[r+8*s] = card
		}
	}
	return deck
}

// New52 liefert ein neues Blatt mit 52 Karten.
func New52() Deck {
	deck := Deck{}
	for _, suit := range []string{"Karo", "Herz", "Pik", "Kreuz"} {
		for _, rank := range []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Bube", "Dame", "König", "Ass"} {
			card := Card{suit, rank}
			deck = append(deck, card)
		}
	}
	return deck
}

// Shuffle mischt das Kartendeck.
func (d Deck) Shuffle() {
	swap := func(i, j int) {
		d[i], d[j] = d[j], d[i]
	}
	rand.Shuffle(len(d), swap)
}

type Hand []Card

// Deal erwartet eine Anzahl an Spielern.
// Liefert für jeden Spieler eine Liste mit 7 Karten.
func (d *Deck) Deal(players int) []Hand {
	hands := []Hand{}
	for i := 0; i < players; i++ {
		h := Hand{}
		for j := 0; j < 7; j++ {
			h = append(h, (*d)[0])
			*d = (*d)[1:]
		}
		hands = append(hands, h)
	}
	return hands
}
