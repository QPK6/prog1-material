package callout

import "fmt"

func ExampleCallout() {
	Callout("Wochenende")

	// Output:
	// wochenende
	// Wochenende
	// WOCHENENDE
	// W   W  OOO   CCCC H   H EEEEE N   N EEEEE N   N DDDD  EEEEE
	// W   W O   O C     H   H E     NN  N E     NN  N D   D E
	// W W W O   O C     HHHHH EEEE  N N N EEEE  N N N D   D EEEE
	// W W W O   O C     H   H E     N  NN E     N  NN D   D E
	//  W W   OOO   CCCC H   H EEEEE N   N EEEEE N   N DDDD  EEEEE
}

func ExampleToUpperFirstLetters() {
	fmt.Println(ToUpperFirstLetters("Bald ist Wochenende"))

	// Output:
	// Bald Ist Wochenende
}

func ExampleAsciiLetter() {
	for _, line := range AsciiLetter("W") {
		fmt.Println(line)
	}

	/* Obige Schleife ist äquivalent zu:
	for i := 0; i < len(AsciiLetter("W")); i++ {
		line := AsciiLetter("W")[i]
		fmt.Println(line)
	}
	*/

	fmt.Println()

	for _, line := range AsciiLetter("N") {
		fmt.Println(line)
	}

	// Output:
	// W   W
	// W   W
	// W W W
	// W W W
	//  W W
	//
	// N   N
	// NN  N
	// N N N
	// N  NN
	// N   N
}
