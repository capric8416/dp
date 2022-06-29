package main

import "fmt"

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	// dp[i][j]中 i表示第i天，j为 [0 - 4] 五个状态，dp[i][j]表示第i天状态j所剩最大现金
	// 0 不做操作
	// 1 第一次买入
	// 2 第一次卖出
	// 3 第二次买入
	// 4 第二次卖出
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 5)
	}

	// dp[0][0] = 0       // 不做操作
	dp[0][1] = -prices[0] // 第一次买入 max(dp[i-1][1], dp[i-1][0]-prices[i]) = -prices[0]
	// dp[0][2] = 0       // 第一次卖出
	dp[0][3] = -prices[0] // 第二次买入 max(dp[i-1][3], dp[i-1][2]-prices[i]) = -prices[0]
	// dp[0][4] = 0       // 第二次卖出

	for i := 1; i < size; i++ {
		prev := (i - 1) % 2
		curr := i % 2
		// 不做操作 = 不做操作，
		dp[curr][0] = dp[prev][0]
		// 第一次买入 = 不做操作，上个状态是不做操作-prices[i]
		dp[curr][1] = max(dp[prev][1], dp[prev][0]-prices[i])
		// 第一次卖出 = 不做操作，上个状态是第一次买入+prices[i]
		dp[curr][2] = max(dp[prev][2], dp[prev][1]+prices[i])
		// 第二次买入 = 不做操作，上个状态是第一次卖出-prices[i]
		dp[curr][3] = max(dp[prev][3], dp[prev][2]-prices[i])
		// 第三次卖出 = 不做操作，上个状态是第二次买入+prices[i]
		dp[curr][4] = max(dp[prev][4], dp[prev][3]+prices[i])
	}

	return dp[(size-1)%2][4]

	// // j为 [0 - 4] 五个状态
	// // 0 不做操作
	// // 1 第一次买入
	// // 2 第一次卖出
	// // 3 第二次买入
	// // 4 第二次卖出
	// dp := make([]int, 5)

	// // dp[0] = 0       // 不做操作
	// dp[1] = -prices[0] // 第一次买入 max(dp[1], dp[0]-prices[i]) = -prices[0]
	// // dp[2] = 0       // 第一次卖出
	// dp[3] = -prices[0] // 第二次买入 max(dp[3], dp[2]-prices[i]) = -prices[0]
	// // dp[4] = 0       // 第二次卖出

	// for i := 1; i < size; i++ {
	// 	// 不做操作 = 不做操作，
	// 	dp[0] = dp[0]
	// 	// 第一次买入 = 不做操作，上个状态是不做操作-prices[i]
	// 	dp[1] = max(dp[1], dp[0]-prices[i])
	// 	// 第一次卖出 = 不做操作，上个状态是第一次买入+prices[i]
	// 	dp[2] = max(dp[2], dp[1]+prices[i])
	// 	// 第二次买入 = 不做操作，上个状态是第一次卖出-prices[i]
	// 	dp[3] = max(dp[3], dp[2]-prices[i])
	// 	// 第三次卖出 = 不做操作，上个状态是第二次买入+prices[i]
	// 	dp[4] = max(dp[4], dp[3]+prices[i])
	// }

	// return dp[4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("6 ==", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println("4 ==", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("0 ==", maxProfit([]int{7, 6, 4, 3, 1}))
}
