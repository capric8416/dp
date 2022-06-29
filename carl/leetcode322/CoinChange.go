package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[j]：凑足总额为j所需钱币的最少个数为dp[j]

	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt
	}

	// 完全背包，但不关心顺序，内外层循环可交换
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt { // 跳过初始化的值
				dp[j] = min(dp[j], dp[j-coins[i]]+1) // 使用了一个零钱，所以加1
			}
		}
	}

	if dp[amount] == math.MaxInt {
		dp[amount] = -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("3 ==", coinChange([]int{1, 2, 5}, 11))
	fmt.Println("-1 ==", coinChange([]int{2}, 3))
	fmt.Println("0 ==", coinChange([]int{1}, 0))
}
