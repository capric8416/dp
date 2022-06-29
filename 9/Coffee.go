package main

import "fmt"

func Coffee(washMachineTime []int, customers int, washTime int, dryTime int) (expected int, got int) {
	expected = CoffeeRecurse(washMachineTime, customers, washTime, dryTime)
	got = CoffeeDp(washMachineTime, customers, washTime, dryTime)
	return
}

func main() {
	fmt.Println(Coffee([]int{3, 1, 7}, 25, 1, 10))
}
