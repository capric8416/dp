package main

import "fmt"

func process1(index int, number int) int {
	if index == number {
		return 1
	}
	if index > number {
		return 0
	}

	return process1(index+1, number) + process1(index+2, number)
}

func jumpFloor1(number int) int {
	// 递归
	// return process1(0, number)

	// 动态规划
	// dp := make([]int, number+1) // 0, 1, ..., number
	// dp[number] = 1              // return 1 if index == number

	// for index := number - 1; index >= 0; index-- {
	// 	dp[index] = dp[index+1]
	// 	if index+2 <= number {
	// 		dp[index] += dp[index+2]
	// 	}
	// }

	// return dp[0]

	// 动态规划, 空间优化
	prevPrev := 1 // prevPrev = 0, prev = 1, current = 1 => prevPrev = prev = 1, prev = current = 1
	prev := 1
	current := 1
	for index := number - 2; index >= 0; index-- {
		current = prev + prevPrev
		prevPrev = prev
		prev = current
	}

	return current
}

func main() {
	fmt.Println(jumpFloor1(2))
	fmt.Println(jumpFloor1(7))
}
