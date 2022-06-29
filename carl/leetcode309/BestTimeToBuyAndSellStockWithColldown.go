package main

import "fmt"

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	// dp[i][j] 表示第i天j状态的最大收益
	// j状态
	// 0 持有股票状态
	// 1 卖出股票状态，两天前就卖出了股票，度过了冷冻期，一直没有操作，保持卖出状态
	// 2 卖出股票状态，今天卖出了股票
	// 3 冷冻期
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0] // 买入
	for i := 1; i < size; i++ {
		prev := (i - 1) % 2
		curr := i % 2
		// 操作一：前一天就是持有股票状态（状态一）
		// 操作二：今天买入了，有两种情况，前一天必须是卖出状态，但卖出后必须冷冻一天，所以
		// 前一天是保持卖出股票状态（状态二）
		// 前一天是冷冻期（状态四），最大收益就是前前天把股票卖出了
		dp[curr][0] = max(dp[prev][0], max(dp[prev][1], dp[prev][3])-prices[i])
		// 达到保持卖出股票状态（状态二），有两个具体操作：
		// 操作一：前一天就是状态二
		// 操作二：前一天是冷冻期（状态四）
		dp[curr][1] = max(dp[prev][1], dp[prev][3])
		// 达到今天就卖出股票状态（状态三），只有一个操作：
		// 操作一：昨天一定是买入股票状态（状态一），今天卖出
		dp[curr][2] = dp[prev][0] + prices[i] // 之前买入状态的收益 + 卖出
		// 达到冷冻期状态（状态四），只有一个操作：
		// 操作一：昨天卖出了股票（状态三）
		dp[curr][3] = dp[prev][2] // 最好的收益是之前把股票卖出了
	}

	index := (size - 1) % 2
	profit := 0
	for state := 1; state <= 3; state++ {
		if dp[index][state] > profit {
			profit = dp[index][state]
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("3 ==", maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println("0 ==", maxProfit([]int{1}))
}
