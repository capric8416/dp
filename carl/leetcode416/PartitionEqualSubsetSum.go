package main

import "fmt"

// 能否把数组分割成和相等的两个子集
func PartitionEqualSubsetSum(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 { // 奇数不可均分为二
		return false
	}

	bagWeight := sum / 2 // 转化为01背包

	dp := make([]int, bagWeight+1) // 背包重量变化范围[0, bagWeight]，因为都是正整数，所以默认初始化为0即可

	for i := 0; i < len(nums); i++ { // 遍历物品
		for j := bagWeight; j >= nums[i]; j-- { // 遍历背包重量
			// dp[i][j] 从[0,i]中任选物品装入重量为j的背包，可以得到的最大价值
			// dp[j] 容量为j的背包，所背的物品价值可以最大为dp[j] -> 套到本题，dp[j]表示背包总容量是j，最大可以凑成j的子集总和为dp[j]
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[bagWeight] == bagWeight // 背包刚好装满物品
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("true ==", PartitionEqualSubsetSum([]int{1, 5, 11, 5}))
	fmt.Println("false ==", PartitionEqualSubsetSum([]int{1, 2, 3, 5}))
}
