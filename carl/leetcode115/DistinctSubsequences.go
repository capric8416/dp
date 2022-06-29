package main

func numDistinct(s string, t string) int {
	length1, length2 := len(s), len(t)

	// dp[i][j] 表示s以i-1结尾，t以j-1结尾，s 的子序列中 t 出现的个数
	dp := make([][]int, length1+1)
	for i := range dp {
		dp[i] = make([]int, length2+1)
		dp[i][0] = 1 // 删除一个字符为空串
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			// s[i - 1] 与 t[j - 1]相等
			// 当s[i - 1] 与 t[j - 1]相等时，dp[i][j]可以有两部分组成。
			// 一部分是用s[i - 1]来匹配，那么个数为dp[i - 1][j - 1]。
			// 一部分是不用s[i - 1]来匹配，个数为dp[i - 1][j]。
			// 这里可能有同学不明白了，为什么还要考虑 不用s[i - 1]来匹配，都相同了指定要匹配啊。
			// 例如： s：bagg 和 t：bag ，s[3] 和 t[2]是相同的，但是字符串s也可以不用s[3]来匹配，即用s[0]s[1]s[2]组成的bag。
			// 当然也可以用s[3]来匹配，即：s[0]s[1]s[3]组成的bag。
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else { // s[i - 1] 与 t[j - 1] 不相等
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[length1][length2]
}
