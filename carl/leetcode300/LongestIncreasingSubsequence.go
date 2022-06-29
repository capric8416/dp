package main

import "fmt"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	// dp[i] 以i结尾的最长递增子序列长度
	// 全部初始化为1
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	length := 1                 // 这两层循环实际上是遍历了所有子串
	for i := 1; i < size; i++ { // 遍历以i结尾
		for j := 0; j < i; j++ { // 遍历不以i结尾
			if nums[i] > nums[j] { // 递推公式 = max(以i结尾，不以i结尾+1)
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > length {
			length = dp[i]
		}
	}

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
