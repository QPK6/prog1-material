package lists

import "fmt"

func ListDemo() {

	// Definiere eine Liste von int.
	// [1 3 5 7 9]
	list1 := []float64{1, 3.5, 5, 7, 9}

	// Gib das erste Element der Liste auf der Konsole aus.
	fmt.Println(list1[0])
	// Gib das letzte Element der Liste auf der Konsole aus.
	fmt.Println(list1[len(list1)-1])

	// Gib ein Element an ungültiger Stelle aus.
	// fmt.Println(list1[5])

	fmt.Println()

	// Gib alle Elemente der Liste aus:
	for i := 0; i < len(list1); i++ {
		fmt.Println(list1[i])
	}
}
