package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	size := len(prices)
	// dp := make([][]int, size) // dp[i][0]表示第i天持有股票所得的最大收益，dp[i][1]表示第i天不持有股票所得的最大收益
	// for i := 0; i < size; i++ {
	// 	dp[i] = make([]int, 2)
	// }

	// dp[0][0] = -prices[0]
	// dp[0][1] = 0
	// for i := 1; i < size; i++ {
	// 	// 维持上一天的状态，或者买入(实际上就是找到更低的买入价格)
	// 	dp[i][0] = max(dp[i-1][0], -prices[i])
	// 	// 维持上一天的状态，或者卖出(上一天状态的收益+卖出的收益)
	// 	dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	// }

	// return dp[size-1][1]

	// // 只依赖了上一行的数据，所以可以用滚动数组优化空间
	// dp := make([][]int, 2) // dp[i][0]表示第i天持有股票所得的最大收益，dp[i][1]表示第i天不持有股票所得的最大收益
	// for i := 0; i < 2; i++ {
	// 	dp[i] = make([]int, 2)
	// }

	// dp[0][0] = -prices[0]
	// dp[0][1] = 0
	// for i := 1; i < size; i++ {
	// 	prev := (i - 1) % 2
	// 	curr := i % 2
	// 	// 维持上一天的状态，或者买入(实际上就是找到更低的买入价格，由于只能买卖一次，所以上一天不持有股票的收益必然是0)
	// 	dp[curr][0] = max(dp[prev][0], -prices[i])
	// 	// 维持上一天的状态，或者卖出(上一天持有股票的收益+卖出的收益)
	// 	dp[curr][1] = max(dp[prev][1], dp[prev][0]+prices[i])
	// }

	// return dp[(size-1)%2][1]

	// 贪心
	low := math.MaxInt
	profit := 0
	for i := 0; i < size; i++ {
		low = min(low, prices[i])           // 求更低的买入价格
		profit = max(profit, prices[i]-low) // 求区间最大收益
		// fmt.Printf("prices[%v]=%v, low=%v, profit=%v\n", i, prices[i], low, profit)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("5 ==", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("0 ==", maxProfit([]int{7, 6, 4, 3, 1}))
}
