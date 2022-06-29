package main

import "math"

/*
 * https://leetcode.com/problems/stickers-to-spell-word/
 *
 * We are given N different types of stickers. Each sticker has a lowercase English word on it.
 *
 * You would like to spell out the given target string by cutting individual letters from your collection of stickers and rearranging them.
 *
 * You can use each sticker more than once if you want, and you have infinite quantities of each sticker.
 *
 * What is the minimum number of stickers that you need to spell out the target? If the task is impossible, return -1.
 *
 * Note: In all test cases, all words were chosen randomly from the 1000 most common US English words, and target was chosen as a concatenation of two random words.
 *
 * Example 1:
 *
 * Input: stickers = ["with","example","science"], target = "thehat"
 * Output: 3
 * Explanation:
 * We can use 2 "with" stickers, and 1 "example" sticker.
 * After cutting and rearrange the letters of those stickers, we can form the target "thehat".
 * Also, this is the minimum number of stickers necessary to form the target string.
 *
 * Example 2:
 *
 * Input: stickers = ["notice","possible"], target = "basicbasic"
 * Output: -1
 * Explanation:
 * We cannot form the target "basicbasic" from cutting letters from the given stickers.
 *
 * Constraints:
 *
 * n == stickers.length
 * 1 <= n <= 50
 * 1 <= stickers[i].length <= 10
 * 1 <= target.length <= 15
 * stickers[i] and target consist of lowercase English letters.
 */
func StikersToSpellWordDp(target string, stickers []string) int {
	if len(target) == 0 || len(stickers) == 0 {
		return 0
	}

	// 由于无法估计 target 字符串变化范围，所以直接上缓存
	cache := make(map[string]int)

	count := processWithCache(cache, target, stickers, "\t")
	if count == math.MaxInt {
		return -1
	}
	return count
}

func processWithCache(cache map[string]int, target string, stickers []string, indent string) (min int) {
	// 为空串了，向上层返回
	if len(target) == 0 {
		min = 0
		return
	}

	if v, ok := cache[target]; ok {
		min = v
		return
	}

	// 表示不含有相同的字符
	min = math.MaxInt

	for i, sticker := range stickers {
		rest := minus(target, sticker, i, indent)
		if len(rest) != len(target) {
			m := process(rest, stickers, indent+"\t")
			min = Min(min, m) // 如果剪除了字符，则表示是可以拼接的
		}
	}

	// 如果可以拼接，证明找到了一张贴纸
	if min != math.MaxInt {
		min += 1
	}
	return
}
