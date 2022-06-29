package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	if n < 0 {
		return -1
	}

	dp := make([]int, n+1) // dp[j]：和为j的完全平方数的最少数量

	// 完全背包，求最小值
	// 初始化: dp[0] = 0, dp[...] = maxint
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	// 完全背包，外层从小到大遍历物品，内层从小到大遍历背包重量
	maxVlue := int(math.Sqrt(float64(n)))
	for i := 0; i <= maxVlue; i++ {
		for j := 0; j <= n; j++ {
			k := j - i*i // 检查索引越界和跳过初始化的maxint值
			if k >= 0 && dp[k] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-i*i]+1) // 选择了i的平方，所以+1
			}
		}
	}

	if dp[n] == math.MaxInt {
		dp[n] = -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("3 ==", numSquares(91))
	fmt.Println("3 ==", numSquares(12))
	fmt.Println("2 ==", numSquares(13))
	fmt.Println("0 ==", numSquares(0))
	fmt.Println("1 ==", numSquares(1))
	fmt.Println("-1 ==", numSquares(-1))
}
