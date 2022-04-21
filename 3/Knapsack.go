package main

import "fmt"

func Knapsack(weights []int, values []int, bagCapacity int) (expected int, got int) {
	expected = KnapsackRecurse(weights, values, bagCapacity)
	got = KnapsackDp(weights, values, bagCapacity)
	return
}

func main() {
	fmt.Println(Knapsack([]int{10, 20, 30}, []int{60, 100, 120}, 51))
}
