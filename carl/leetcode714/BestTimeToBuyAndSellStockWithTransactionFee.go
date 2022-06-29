package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	// dp[i][0] 持有股票
	// dp[i][1] 不持有股票
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] - fee
	for i := 1; i < size; i++ {
		prev := (i - 1) % 2
		curr := i % 2

		dp[curr][0] = max(dp[prev][0], dp[prev][1]-prices[i]-fee)
		dp[curr][1] = max(dp[prev][1], dp[prev][0]+prices[i])
	}

	return dp[(size-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("8 ==", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println("6 ==", maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}
