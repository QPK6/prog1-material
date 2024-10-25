package main

import (
	"fmt"
	"math/rand"
)

// ReadNumber fragt den Benutzer nach einer Zahl und liefert diese zurück.
func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	expected_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, expected_number) {
			fmt.Println("Richtig geraten :-)")
			return
		} else if guess < expected_number {
			fmt.Println("Die gesuchte Zahl ist größer.")
		} else {
			fmt.Println("Die gesuchte Zahl ist kleiner.")
		}
	}
	fmt.Println("Zu viele falsche Zahlen :-(")
}

// Prüft, ob die geratene Zahl der erwarteten entspricht.
func NumberGood(guess, expected int) bool {
	return guess == expected
}

func main() {
	GuessingGame()
}
