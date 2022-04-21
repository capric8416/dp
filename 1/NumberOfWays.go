package main

import "fmt"

func NumberOfWays(maxPos, currentPos, moveSteps, targetPos int) (expected int, got int) {
	got = NumberOfWaysDp(maxPos, currentPos, moveSteps, targetPos)
	expected = NumberOfWaysRecurse(maxPos, currentPos, moveSteps, targetPos)
	return expected, got
}

func main() {
	fmt.Println(NumberOfWays(7, 2, 4, 4))
}
