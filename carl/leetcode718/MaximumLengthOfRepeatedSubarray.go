package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return 0
	}

	// dp[i][j] 表示nums1以i-1结尾, nums2以j-1结尾公共子数组最大长度
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 首行首列都是0
	length := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > length {
					length = dp[i][j]
				}
			}
		}
	}

	// 有问题，等待处理
	// // 滚动数组，内循环从大到小遍历，防止覆盖
	// dp := make([]int, m+1)
	// length := 0
	// for i := m - 1; i >= 0; i-- {
	// 	for j := n - 1; j >= 0; j-- {
	// 		if nums1[i] == nums2[j] {
	// 			dp[j] = dp[j+1] + 1
	// 			if dp[j] > length {
	// 				length = dp[j]
	// 			}
	// 		} else {
	// 			dp[i] = 0
	// 		}
	// 	}
	// }

	return length
}
func main() {
	fmt.Println(findLength([]int{1, 0, 0, 0, 1}, []int{1, 0, 0, 1, 1}))
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
