package main

import "fmt"

func LongestPalindromicSubsequence(s string) (expected int, got int) {
	expected = LongestPalindromicSubsequenceRecurse(s)
	got = LongestPalindromicSubsequenceDp(s)
	return
}
func main() {
	fmt.Println(LongestPalindromicSubsequence("a12b3c43def2ghi1kpm"))
}
