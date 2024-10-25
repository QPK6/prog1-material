package findlist

import "fmt"

func ExampleFindElementLinear() {
	list := []int{1, 3, 5, 7, 9}

	fmt.Println(FindElementLinear(list, 5))

	// Output:
	// 2
}

func ExampleFindElementBinary() {
	list := []int{1, 3, 5, 7, 9}

	fmt.Println(FindElementBinary(list, 5))
	fmt.Println(FindElementBinary(list, 9))
	fmt.Println(FindElementBinary(list, -3))
	fmt.Println(FindElementBinary(list, 42))

	// Output:
	// 2
	// 4
	// -1
	// -1
}
