package main

/*
 * 516. Longest Palindromic Subsequence
 *
 * Given a string s, find the longest palindromic subsequence's length in s.
 *
 * A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
 *
 *
 *
 * Example 1:
 *
 * Input: s = "bbbab"
 * Output: 4
 * Explanation: One possible longest palindromic subsequence is "bbbb".
 * Example 2:
 *
 * Input: s = "cbbd"
 * Output: 2
 * Explanation: One possible longest palindromic subsequence is "bb".
 *
 *
 * Constraints:
 *
 * 1 <= s.length <= 1000
 * s consists only of lowercase English letters.
 */

func LongestPalindromicSubsequenceDp(s string) int {
	if len(s) == 0 {
		return 0
	}

	size := len(s)
	dp := make([][]int, size)
	for left := 0; left < size; left++ {
		dp[left] = make([]int, size)
		// left == right 时, 只有一个字符, 对角线初始化为0
		dp[left][left] = 1
		// left == right-1 时, 只有两个字符, 如果相同初始化为2, 如果不同初始化为1
		if right := left + 1; right < size {
			if s[left] == s[right] {
				dp[left][right] = 2
			} else {
				dp[left][right] = 1
			}
		}
	}

	// left <= right, 所以对角线下方无需初始化
	for left := size - 3; left >= 0; left-- {
		for right := left + 2; right < size; right++ {
			c1 := dp[left+1][right-1]
			c2 := dp[left+1][right]
			c3 := dp[left][right-1]
			c4 := 0
			if s[left] == s[right] {
				c4 = 2 + c1
			}
			// 优化:
			// 左边格子, 下方格式都大于左下
			// 所以可以去掉 c1 分支, 变成
			// dp[left][right] = Max(dp[left+1][right], dp[left][right-1])
			// if s[left] == s[right] {
			// 	dp[left][right] = Max(dp[left][right], 2+dp[left+1][right-1])
			// }

			dp[left][right] = Max(Max(c1, c2), Max(c3, c4))
		}
	}

	return dp[0][size-1]
}
