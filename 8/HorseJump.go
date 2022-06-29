package main

import "fmt"

func HourseJump(targetXPos int, targetYPos int, steps int) (expected int, got int) {
	expected = HourseJumpRecurse(targetXPos, targetYPos, steps)
	got = HourseJumpDp(targetXPos, targetYPos, steps)
	return
}

func main() {
	fmt.Println(HourseJump(2, 1, 3))
}
