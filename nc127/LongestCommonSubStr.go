package main

import "fmt"

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	// write code here

	str1Length := len(str1)
	str2Length := len(str2)

	dp := make([][]int, str1Length)
	for i := 0; i < str1Length; i++ {
		dp[i] = make([]int, str2Length)
	}

	for i, j := 0, 0; i < str1Length; i++ {
		if str1[i] == str2[j] {
			dp[i][j] = 1
		}
	}
	for i, j := 0, 0; j < str2Length; j++ {
		if str1[i] == str2[j] {
			dp[i][j] = 1
		}
	}

	lastOffset, maxLength := 0, 0
	for i := 1; i < str1Length; i++ {
		for j := 1; j < str2Length; j++ {
			if str1[i] == str2[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if dp[i][j] > maxLength {
					lastOffset = i
					maxLength = dp[i][j]
				}
			}
		}
	}

	return str1[lastOffset-maxLength+1 : lastOffset+1]
	// return dp[str1Length-1][str2Length-1]
	// return process(str1, str2, str1Length-1, str2Length-1)
}

// 样本对应模型
func process(str1 string, str2 string, i int, j int, max *int) int {
	// 二维dp
	// i 变化范围是[0, str1Length-1]
	// j 变化范围是[0, str2Length-1]
	// 初始化 dp[0][0] = str1[0] == str2[0] ? 1 : 0

	if i == 0 && j == 0 { // 都只剩下一个字符了
		if str1[i] == str2[j] {
			return 1
		}
		return 0
	} else if i == 0 { // str1 只剩下一个字符了
		if str1[i] == str2[j] {
			return 1
		}
		return 0 // 因为子串是连续的，所以不相等不停止继续比较了
		// return process(str1, str2, i, j-1)
	} else if j == 0 { // str2 只剩下一个字符了
		if str1[i] == str2[j] {
			return 1
		}
		return 0 // 因为子串是连续的，所以不相等不停止继续比较了
		// return process(str1, str2, i-1, j)
	} else {
		if str1[i] == str2[j] {
			c := 1 + process(str1, str2, i-1, j-1, max)
			if c > *max {
				*max = c
			}
		}
		return 0 // 因为子串是连续的，所以不相等不停止继续比较了
	}
}

func main() {
	fmt.Println(LCS("abcdefg", "ab1cdefgabc1defg"))
	fmt.Println(LCS("22222", "22222"))
	fmt.Println(LCS("1AB2345CD", "12345EF"))
}
