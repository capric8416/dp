package main

import "fmt"

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size) // 从[0-i]间房子中最多可以偷到多少钱
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < size; i++ {
		// 不偷第i个房子
		// 偷第i个房子，即不偷i-1个房子(偷第i-2个房子和第i个房子)
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("4 ==", rob([]int{1, 2, 3, 1}))
	fmt.Println("12 ==", rob([]int{2, 7, 9, 3, 1}))
}
