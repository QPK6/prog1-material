package main

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	my_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Richtig geraten :-)")
			return
		} else {
			fmt.Println("Falsch geraten (Muhaha)")
		}
	}
	fmt.Println("Zu viele falsche Zahlen :-(")
}

func NumberGood(g, e int) bool {
	return g == e
}

func main() {
	GuessingGame()
}
