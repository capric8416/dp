package main

import "fmt"

func change(amount int, coins []int) int {
	// 组合
	// dp[j] 表示金额为i的背包，有多少种方法可以凑成 (零钱数组coins，每种面额零钱有无数种)
	dp := make([]int, amount+1)

	// 组合：外层遍历物品，内层遍历背包
	// 排列：外层遍历背包，内层遍历物品
	dp[0] = 1 // 凑成总金额0的货币组合数为1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 组合
			dp[j] += dp[j-coins[i]] // dp[j] （考虑coins[i]的组合总和） 就是所有的dp[j - coins[i]]（不考虑coins[i]）相加
		}
	}

	return dp[amount]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("5 ==", change(5, []int{1, 2, 5}))
	fmt.Println("0 ==", change(3, []int{2}))
	fmt.Println("1 ==", change(10, []int{10}))
}
