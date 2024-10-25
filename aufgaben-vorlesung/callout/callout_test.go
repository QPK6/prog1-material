package callout

func ExampleCallout() {
	Callout("Wochenende")

	// Output:
	// wochenende
	// Wochenende
	// WOCHENENDE
	// W   W  OOO   CCC H   H  EEEEE N   N EEEEE N   N DDDD  EEEEE
	// W   W O   O C    H   H  E     NN  N E     NN  N D   D E
	// W W W O   O C    HHHHH  EEEE  N N N EEEE  N N N D   D EEEE
	// W W W O   O C    H   H  E     N  NN E     N  NN D   D E
	//  W W   OOO   CCC H   H  EEEEE N   N EEEEE N   N DDDD  EEEEE
}
