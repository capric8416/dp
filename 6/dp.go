package main

/*
 * 1143. Longest Common Subsequence
 *
 * Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
 *
 * A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
 *
 * For example, "ace" is a subsequence of "abcde".
 * A common subsequence of two strings is a subsequence that is common to both strings.
 *
 *
 *
 * Example 1:
 *
 * Input: text1 = "abcde", text2 = "ace"
 * Output: 3
 * Explanation: The longest common subsequence is "ace" and its length is 3.
 * Example 2:
 *
 * Input: text1 = "abc", text2 = "abc"
 * Output: 3
 * Explanation: The longest common subsequence is "abc" and its length is 3.
 * Example 3:
 *
 * Input: text1 = "abc", text2 = "def"
 * Output: 0
 * Explanation: There is no such common subsequence, so the result is 0.
 *
 *
 * Constraints:
 *
 * 1 <= text1.length, text2.length <= 1000
 * text1 and text2 consist of only lowercase English characters.
 */

func LongestCommonSubsequenceDp(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1))
	for r := 0; r < len(text1); r++ {
		dp[r] = make([]int, len(text2))
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	}
	for j := 1; j < len(text2); j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(text1); i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			c1 := dp[i-1][j]
			c2 := dp[i][j-1]
			c3 := 0
			if text1[i] == text2[j] {
				c3 = 1 + dp[i-1][j-1]
			}
			dp[i][j] = Max(c1, Max(c2, c3))
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}
