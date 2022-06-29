package main

func longestCommonSubsequence(text1 string, text2 string) int {
	length1, lenght2 := len(text1), len(text2)
	if length1 == 0 || lenght2 == 0 {
		return 0
	}

	// dp[i][j] 表示 text1以i-1结尾， text2以j-1结尾的公共子序列最大长度
	dp := make([][]int, length1+1)
	for i := range dp {
		dp[i] = make([]int, lenght2+1)
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= lenght2; j++ {
			if text1[i-1] == text2[j-1] { // i-1, j-1结尾
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// i-2, j-1结尾
				// i-1, j-2结尾
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
