package main

import "fmt"

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	count := 2 * k

	// 买入然后卖出称为一次操作，所以k次交易实际上是指k次买入和卖出
	// 状态j
	// 0    不做操作
	// 1    第一次买入
	// 2    第一次卖出
	// ...
	// 2k-1 最后一次买入
	// 2k   最后一次卖出
	dp := make([][]int, 2) // dp[i][j] 第i天j状态的最大收益
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, count+1)
	}

	for j := 1; j <= count; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < size; i++ {
		prev := (i - 1) % 2
		curr := i % 2

		dp[curr][0] = dp[prev][0]
		for j := 1; j <= count; j += 2 {
			// 买入 = 不做操作 or 上一次卖出 - 当前价格(买入)
			dp[curr][j] = max(dp[prev][j], dp[prev][j-1]-prices[i])
			// 卖出 = 不做操作 or 上一次买入 + 当前价格(卖出)
			dp[curr][j+1] = max(dp[prev][j+1], dp[prev][j]+prices[i])
		}
	}

	return dp[(size-1)%2][count]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("2 ==", maxProfit(2, []int{2, 4, 1}))
	fmt.Println("7 ==", maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
