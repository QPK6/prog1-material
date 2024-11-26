package pointer

import "fmt"

func Example_pointer() {
	var x int

	// Die Adresse von x in einem int-Pointer px speichern.
	var px *int = &x

	// Speicheradresse von x ausgeben.
	fmt.Println(px)

	// Den Wert von x ausgeben, aber durch den Pointer hindurch.
	fmt.Println(*px)

	// x verändern
	x = 42

	fmt.Println(*px)

	// Den Wert hinter px verändern.
	*px = 23

	fmt.Println(x)

	// Output:
}

// Add2 addiert 2 auf den Wert von x
func Add2(elefant *int) {
	*elefant = *elefant + 2
	fmt.Println(*elefant)
}

// Mult multipliziert den Wert von x mit 2
func Mult2(elefant *int) {
	*elefant = *elefant * 2
	fmt.Println(*elefant)
}

func Example_call_by_ref() {
	x := 40
	y := 17
	Add2(&x)
	Mult2(&y)

	fmt.Println(x)

	// Output:
	// 42
}
