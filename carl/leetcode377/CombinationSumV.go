package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	// 最值型完全背包，内外层循环顺序无关紧要
	// 组合数：外层遍历物品，内层遍历背包
	// 排列数：外层遍历背包，内层遍历物品
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Println("7 ==", combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println("0 ==", combinationSum4([]int{9}, 1))
}
