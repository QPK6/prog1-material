package ean

import "fmt"

func ExampleInputOk() {
	fmt.Println(InputOk("1231231231231"))
	fmt.Println(InputOk("12312312"))
	fmt.Println(InputOk("123123123123A"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsDigit() {
	fmt.Println(IsDigit('0'))
	fmt.Println(IsDigit('9'))
	fmt.Println(IsDigit('a'))

	// Output:
	// true
	// true
	// false
}

func ExampleToIntList() {
	fmt.Println(ToIntList("1234"))

	// Output:
	// [1 2 3 4]
}

func ExampleCheckSum() {
	fmt.Println(CheckSum([]int{9, 0, 9, 9, 9, 9, 9, 5, 4, 3, 2, 1, 7}))
	fmt.Println(CheckSum([]int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0, 3}))

	// Output:
	// 0
	// 0
}

func ExampleEanOk() {
	fmt.Println(EanOk("9099999543217"))
	fmt.Println(EanOk("9780345391803"))
	fmt.Println(EanOk("9780345391809"))
	fmt.Println(EanOk("978034539180"))
	fmt.Println(EanOk("978034539180A"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}
