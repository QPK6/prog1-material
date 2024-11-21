package callout

import (
	"fmt"
	"strings"
)

// Callout erwartet einen String und gibt ihn vier Mal
// in aufsteigender "Lautstärke" auf der Konsole aus.
// Die erste Ausgabe erfolgt komplett in Kleinbuchstaben,
// die zweite mit großen Anfangsbuchstaben,
// die dritte vollständig in Großbuchstaben und
// die vierte als große Ascii-Art-Schrift.
func Callout(s string) {
	fmt.Println(strings.ToLower(s))
	fmt.Println(ToUpperFirstLetters(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(ToAsciiArt(s))
}

// s in Kleinbuchstaben umwandeln,
// dann jeweils ersten Buchstaben eines Wortes
// in Großbuchstaben umwandeln.
func ToUpperFirstLetters(s string) string {

	// Den String in eine Liste von Wörtern umwandeln.
	words := strings.Split(s, " ")

	// Die ersten Buchstaben der Wörter in Großbuchstaben umwandeln.
	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		firstLetter := currentWord[0]
		if firstLetter >= 'a' {
			newFirstLetter := firstLetter - ('a' - 'A') // 'a' - 'A' == 32
			newWord := string(newFirstLetter) + currentWord[1:]
			words[i] = newWord
		}
	}

	return strings.Join(words, " ")
}

// s in AsciiArt umwandeln
func ToAsciiArt(s string) string {

	lines := []string{"", "", "", "", ""}

	// Durch s laufen und für jeden Buchstaben
	// die Strings bestimmen, die dem Buchstaben entsprechen.
	for pos, c := range s {
		char_lines := AsciiLetter(string(c))

		// Diese String-Listen an lines anhängen.
		for i := 0; i < 5; i++ {
			lines[i] += char_lines[i]
			if pos < len(s)-1 {
				lines[i] += " "
			}
		}
	}
	return strings.Join(lines, "\n")
}

func AsciiLetter(c string) []string {
	switch strings.ToUpper(c) {
	case "W":
		return []string{
			"W   W",
			"W   W",
			"W W W",
			"W W W",
			" W W "}

	case "O":
		return []string{
			" OOO ",
			"O   O",
			"O   O",
			"O   O",
			" OOO ",
		}
	case "C":
		return []string{
			" CCCC",
			"C    ",
			"C    ",
			"C    ",
			" CCCC",
		}
	case "H":
		return []string{
			"H   H",
			"H   H",
			"HHHHH",
			"H   H",
			"H   H",
		}
	case "E":
		return []string{
			"EEEEE",
			"E    ",
			"EEEE ",
			"E    ",
			"EEEEE",
		}
	case "N":
		return []string{
			"N   N",
			"NN  N",
			"N N N",
			"N  NN",
			"N   N",
		}
	case "D":
		return []string{
			"DDDD ",
			"D   D",
			"D   D",
			"D   D",
			"DDDD ",
		}
	default:
		return []string{}
	}
}
