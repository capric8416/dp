package main

func isSubsequence(s string, t string) bool {
	length1, length2 := len(s), len(t)
	if (length1 == 0 && length2 == 0) || (length1 == 0 && length2 != 0) {
		return true
	}

	dp := make([][]int, length1+1)
	for i := range dp {
		dp[i] = make([]int, length2+1)
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[length1][length2] == length1
}
