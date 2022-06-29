package main

import "fmt"

/**
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	// write code here

	// dp := make([]int, n+1)
	// dp[1] = 1
	// if n >= 2 {
	// 	dp[2] = 1
	// }
	prevPrev := 1
	prev := 1
	current := 1

	for index := 3; index <= n; index++ {
		// dp[index] = dp[index-1] + dp[index-2]
		current = prevPrev + prev
		prevPrev = prev
		prev = current
	}
	return current
	// return dp[n]
	// return process(n, n)
}

func process(current int, n int) int {
	if current == 1 || current == 2 {
		return 1
	}

	return process(current-1, n) + process(current-2, n)
}

func main() {
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
}
