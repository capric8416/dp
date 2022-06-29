package main

import "fmt"

func LongestCommonSubsequence(text1 string, text2 string) (expected int, got int) {
	expected = LongestCommonSubsequenceRecurse(text1, text2)
	got = LongestCommonSubsequenceDp(text1, text2)
	return
}

func main() {
	fmt.Println(LongestCommonSubsequence("abcde", "ace"))
	fmt.Println(LongestCommonSubsequence("abc", "abc"))
	fmt.Println(LongestCommonSubsequence("abc", "def"))
	fmt.Println(LongestCommonSubsequence("kabcfffffffffdm", "agcyfxdomu"))
}
