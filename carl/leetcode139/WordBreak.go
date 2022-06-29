package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// 完全背包，求可行性
	length := len(s)
	dp := make([]bool, length+1) // 字符串s长度为j的话，dp[j]为true，表示可以拆分为一个或多个在字典中出现的单词
	dp[0] = true                 // dp[0]初始为true

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// 完全背包，内外层顺序无关紧要
	// 但这里要求子串，所以外层遍历背包，内层遍历物品
	// 这里的两层循环，也就是遍历了s的所有子串
	for j := 1; j <= length; j++ {
		for i := 0; i < j; i++ {
			// 如果子串s[i:j]出现在字典里，即可拆分
			// 那么子串前面的部分s[:i]也应该是可以拆分的，应该的状态就是dp[i]
			if dp[i] && wordMap[s[i:j]] {
				dp[j] = true
			}
		}
	}

	return dp[length]
}

func main() {
	fmt.Println("true ==", wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println("true ==", wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println("false ==", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println("true ==", wordBreak("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}))
}
