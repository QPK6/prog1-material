package pointer

import "fmt"

type Uhrzeit struct {
	Stunde  int
	Minute  int
	Sekunde int
}

// Tick zählt die Sekunde um 1 hoch.
func Tick(u *Uhrzeit) {
	u.Sekunde++
}

// Tack zählt die Sekunde um 1 hoch.
func (u *Uhrzeit) Tack() {
	u.Sekunde++
	// if u.Sekunde >= 60 {
	// 	u.Sekunde = 0
	// 	u.Minute++
	// }
	// if u.Minute >= 60 {
	// 	u.Minute = 0
	// 	u.Stunde++
	// }
	// if u.Stunde >= 24 {
	// 	u.Stunde = 0
	// }
	u.Minute += u.Sekunde / 60
	u.Stunde += u.Minute / 60
	u.Stunde = u.Stunde % 24
	u.Minute %= 60
	u.Sekunde %= 60
}

// Tock zählt auch die Sekunde hoch und liefert das Ergebnis als neues Struct.
func Tock(u Uhrzeit) Uhrzeit {
	u.Sekunde++
	return u
}

func ExampleTick() {
	u := Uhrzeit{0, 0, 0}

	Tick(&u)
	u = Tock(u)
	u.Tack()

	fmt.Println(u)

	// Output:

}

func ExampleUhrzeit_Tack_overflow() {
	u1 := Uhrzeit{0, 0, 59}
	u2 := Uhrzeit{0, 59, 59}
	u3 := Uhrzeit{23, 59, 59}

	u1.Tack()
	u2.Tack()
	u3.Tack()

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)

	// Output:
	// {0 1 0}
	// {1 0 0}
	// {0 0 0}
}
