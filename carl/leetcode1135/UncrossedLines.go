package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	length1, lenght2 := len(nums1), len(nums2)
	if length1 == 0 || lenght2 == 0 {
		return 0
	}

	dp := make([][]int, length1+1)
	for i := range dp {
		dp[i] = make([]int, lenght2+1)
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= lenght2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[length1][lenght2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
