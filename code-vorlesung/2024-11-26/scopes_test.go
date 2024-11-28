package scopes

import "fmt"

func Example_scopes() {
	x := 1
	{
		x := 2
		{
			x := 3
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)

	//Output:
}

func Example_scopes_loop() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(m[i][j])
		}
	}
	for i := 0; i < 3; i++ { // für zwei Schleifen die gleichen Zähler -> Scopes
		for i := 0; i < 3; i++ {
			fmt.Println(m[i][i])
		}
	}
	//Output:
}

func Foo(x int) {
	fmt.Println(x * 4)
	Bar(x + 2)
}

func Bar(x int) { //x==44
	Baz(x / 2)
	fmt.Println(x)
}

func Baz(x int) {
	x += 17
	fmt.Println(x)
}

func EcampleFoo() {
	Foo(42)
	//Output:
	//168
	//39
	//44
}

func Fool(x int) int {
	if x == 0 {
		return 1
	}
	return 2 * Fool(x-1)
}

func Example_scopes_rec() {
	fmt.Println(Fool(0))
	fmt.Println(Fool(1))
	fmt.Println(Fool(2))
	fmt.Println(Fool(3))
	//Output:
	//1
	//2
	//4
	//8
}
