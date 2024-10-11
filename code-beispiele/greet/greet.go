package greet

import "fmt"

func Greet(name string) string { // (name string) = Zeichenkette -> string erwartete Zeichenkette
	return fmt.Sprintf("Hello %s!", name) // print formate %s = Platzhalter %d=dezimal %c=Zeichen %i=integer
}
