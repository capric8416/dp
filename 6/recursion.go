package main

func LongestCommonSubsequenceRecurse(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	return process(text1, text2, len(text1)-1, len(text2)-1)
}

// text1[0...i] 与 text2[0...j] 的最长公共子序列
func process(text1 string, text2 string, i int, j int) int {
	// text1 和 text2 都只剩下一个字符了，如果相等返回1，否则返回0
	if i == 0 && j == 0 {
		if text1[i] == text2[j] {
			return 1
		}
		return 0
	} else if i == 0 {
		// text1 只剩下一个字符了，如果相等直接返回1，否则继续查找 text2 剩下的字符里面是否和相同的字符
		if text1[i] == text2[j] {
			return 1
		}
		return process(text1, text2, i, j-1)
	} else if j == 0 {
		// text2 只剩下一个字符了，如果相等直接返回1，否则继续查找 text1 剩下的字符里面是否和相同的字符
		if text1[i] == text2[j] {
			return 1
		}
		return process(text1, text2, i-1, j)
	} else {
		// 样本对应模型，针对结尾字符考虑以下3种情形
		// 1) 不考虑以 i 结尾，但可能以 j 结尾
		// 2) 可能以 i 结尾，但不考虑以 j 结尾
		// 3) 可能以 i 结尾，也可能以 j 结尾
		// 虽然3种情形有交集，但因为是求最长公共子序列，所以没关系
		c1 := process(text1, text2, i-1, j)
		c2 := process(text1, text2, i, j-1)
		c3 := 0
		if text1[i] == text2[j] {
			c3 = 1 + process(text1, text2, i-1, j-1)
		}
		return Max(c1, Max(c2, c3))
	}
}
