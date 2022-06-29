package main

import "fmt"

func ConvertToLetterString(digits string) (expected int, got int) {
	expected = ConvertToLetterStringRecurse(digits)
	got = ConvertToLetterStringDp(digits)
	return
}

func main() {
	fmt.Println(ConvertToLetterString("0"))
	fmt.Println(ConvertToLetterString("1"))
	fmt.Println(ConvertToLetterString("101"))
	fmt.Println(ConvertToLetterString("111"))
	fmt.Println(ConvertToLetterString("305"))
}
