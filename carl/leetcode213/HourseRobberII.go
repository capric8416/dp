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

	// 偷第一个，不偷最后一个
	// 不偷第一个，偷最后一个
	return max(robPartition(nums, 0, size-2), robPartition(nums, 1, size-1))
}

func robPartition(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	size := right - left + 1
	dp := make([]int, size) // 从[left, right]间房子中最多可以偷到多少钱
	dp[0] = nums[left]
	dp[1] = max(nums[left], nums[left+1])

	for i := 2; i < size; i++ {
		// 不偷第i个房子
		// 偷第i个房子，即不偷i-1个房子(偷第i-2个房子和第i个房子)
		dp[i] = max(dp[i-1], dp[i-2]+nums[left+i])
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
	fmt.Println("3 ==", rob([]int{2, 3, 2}))
	fmt.Println("4 ==", rob([]int{1, 2, 3, 1}))
	fmt.Println("3 ==", rob([]int{1, 2, 3}))
}
