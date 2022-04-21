package main

import "fmt"

func CardsInLine(arr []int) (got int, expected int) {
	got = CardsInLineDp(arr, 0, len(arr)-1)
	expected = CardsInLineRecurse(arr, 0, len(arr)-1)
	return got, expected
}

func main() {
	fmt.Println(CardsInLine([]int{5, 20, 15, 10, 20}))
	fmt.Println(CardsInLine([]int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}))
}
