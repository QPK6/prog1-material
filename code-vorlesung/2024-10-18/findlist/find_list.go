package findlist

// FindElementLinear erwartet eine Liste von Zahlen und eine weitere Zahl n.
// Die Funktion liefert die Position von n in der Liste.
// Falls n nicht in der Liste vorkommt, liefert die Funktion -1.
//
// Die Funktion führt eine lineare Suche in der Liste durch.
// D.h. sie durchsucht die Liste von Anfang bis Ende.
func FindElementLinear(list []int, n int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == n {
			return i
		}
	}
	return -1
}

// FindElementBinary erwartet eine Liste von Zahlen und eine weitere Zahl n.
// Die Funktion liefert die Position von n in der Liste.
// Falls n nicht in der Liste vorkommt, liefert die Funktion -1.
//
// Die Funktion führt eine binäre Suche in der Liste durch.
// D.h. sie startet in der Mitte und macht dann im linken oder rechten Teil weiter.
// Voraussetzung: Die Liste muss sortiert sein!
func FindElementBinary(list []int, n int) int {

	if len(list) == 0 {
		return -1
	}

	m := len(list) / 2
	if list[m] == n {
		return m
	}

	if n < list[m] {
		return FindElementBinary(list[:m], n)
	}

	result := FindElementBinary(list[m+1:], n)
	if result == -1 {
		return -1
	}
	return result + m + 1
}
