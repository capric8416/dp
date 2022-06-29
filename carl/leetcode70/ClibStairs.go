package main

import "fmt"

/*
01背包
最大值，外层从小到大遍历物品，内层从大到小遍历背包                              dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
最大值，外层从小到大遍历物品，内层从大到小遍历背包 dp[0] = 0, dp[...] = maxint, dp[j] = min(dp[j], dp[j-weight[i]]+value[i]) if j - weight[i] >= 0 && dp[j-weight[i]] != maxint
组合数，外层从小到遍历物品，内层从大到小遍历背包   dp[0] = 1,                   dp[j] += dp[j-weight[i]]

完全背包
最大值，外层从小到大遍历物品，内层从小到大遍历背包，反之亦可                              dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
最小值，外层从小到大遍历物品，内层从小到大遍历背包，反之亦可 dp[0] = 0, dp[...] = maxint, dp[j] = min(dp[j], dp[j-weight[i]]+value[i])  if j - weight[i] >= 0 && dp[j-weight[i]] != maxint
组合数，外层从小到大遍历物品，内层从大到小遍历背包          dp[0] = 1,                   dp[j] += dp[j-weight[i]]
排列数，外层从小到大遍历背包，内层从小到大遍历物品          dp[0] = 1,                   dp[j] += dp[j-weight[i]]
*/

func climbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	// 先爬一个台阶，后爬两个台阶；跟先爬两个台阶，再爬一个台阶；顺序不一样，所以是排列
	// 排列：外层从小到大遍历背包，内层从小到大遍历物品
	for j := 0; j <= n; j++ {
		for i := 1; i <= 2; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println("2 ==", climbStairs(2))
	fmt.Println("3 ==", climbStairs(3))
}
