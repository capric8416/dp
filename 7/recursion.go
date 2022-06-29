package main

// 方法1: 将 s 反序成 s',求二者的最长公共子序列
func LongestPalindromicSubsequenceRecurse(s string) int {
	if len(s) == 0 {
		return 0
	}

	return process(s, 0, len(s)-1)
}

func process(s string, left int, right int) int {
	if left == right {
		// 只有一个字符, 自己和自己肯定可以构成回文
		return 1
	}
	if left == right-1 {
		// 如果只有两个字符, 并且相同, 那回文长度为2
		if s[left] == s[right] {
			return 2
		}
		// 要么左边的字符构成回文, 要么右边字符构成回文
		return 1
	}

	// 范围尝试模型, 讨论开头和结尾4种情形 (样本对应模型, 讨论结尾3种情形)
	// 1) 不以 left 开头, 也不以 right 结尾
	// 2) 以 left 开头, 不以 right 结尾
	// 3) 不以 left 开头, 不以 right 结尾
	// 4) 以 left 开头, 也以 right 结尾
	// 求4种情形的大值
	c1 := process(s, left+1, right-1)
	c2 := process(s, left+1, right)
	c3 := process(s, left, right-1)
	c4 := 0
	if s[left] == s[right] {
		c4 = 2 + c1
	}

	return Max(Max(c1, c2), Max(c3, c4))
}
