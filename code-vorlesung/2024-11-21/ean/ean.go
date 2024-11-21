package ean

import "unicode"

// InputOk prüft, ob code 13 Stellen hat
// und nur aus Ziffern besteht.
func InputOk(code string) bool {
	if len(code) != 13 {
		return false
	}

	for _, c := range code {
		if !unicode.IsDigit(c) { // aus Paket unicode, eigene Funktion unnötig.
			return false
		}
	}
	return true
}

// IsDigit erwartet ein Zeichen c und prüft,
// ob dieses eine Ziffer ist.
func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// ToIntList erwartet einen String "code",
// der nur aus Ziffern besteht.
// Liefert eine Liste dieser Ziffern als []int.
func ToIntList(code string) []int {
	result := []int{}
	//result := make([]int, len(code))

	for i := 0; i < len(code); i++ {
		//result[i] = int(code[i] - '0')
		result = append(result, int(code[i]-'0'))
	}

	return result
}

// CheckSum erwartet eine Liste von Ziffern und
// berechnet deren EAN-Prüfsumme.
func CheckSum(digits []int) int {
	result := 0

	// for i := 0; i < len(digits); i++ {
	// 	n := digits[i]
	for i, n := range digits {
		//result += (2*(i%2) + 1) * n // ginge auch, ist nur nicht lesbar
		if i%2 == 0 {
			result += n
		} else {
			result += 3 * n
		}
	}
	return (10 - (result % 10)) % 10
}

// EanOk erwartet einen String "code" und liefert true,
// genau dann wenn code eine gültige EAN-Nummer ist.
func EanOk(code string) bool {
	if !InputOk(code) {
		return false
	}

	digits := ToIntList(code)

	checksum := CheckSum(digits)

	return checksum == 0
}
