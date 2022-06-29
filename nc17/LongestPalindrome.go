package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param A string字符串
 * @return int整型
 */
func getLongestPalindrome(A string) int {
	// write code here

	length := len(A)
	if length < 2 {
		return length
	}

	dp := make([][]bool, length)
	for left := 0; left < length; left++ {
		dp[left] = make([]bool, length)
	}

	maxLength := 1
	for right := 1; right < length; right++ {
		for left := 0; left <= right; left++ {
			if A[left] != A[right] {
				continue
			}

			if left == right {
				// "a"
				dp[left][right] = true
			} else if right-left <= 2 {
				// "aa", "aba"
				dp[left][right] = true
			} else {
				// "a*************a"
				dp[left][right] = dp[left+1][right-1]
			}

			if dp[left][right] && right-left+1 > maxLength {
				maxLength = right - left + 1
			}
		}
	}

	return maxLength
}

func main() {
	fmt.Println(getLongestPalindrome("ccbcabaabba"))
	fmt.Println(getLongestPalindrome("ababc"))
	fmt.Println(getLongestPalindrome("abbba"))
	fmt.Println(getLongestPalindrome("b"))
}
