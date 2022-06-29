package main

import (
	"math"
)

func StikersToSpellWordRecurse(target string, stickers []string) int {
	if len(target) == 0 || len(stickers) == 0 {
		return 0
	}

	// fmt.Printf("0 = %#v %v\n", target, stickers)

	count := process(target, stickers, "\t")
	if count == math.MaxInt {
		return -1
	}
	return count
}

func process(target string, stickers []string, indent string) (min int) {
	// 为空串了，向上层返回
	if len(target) == 0 {
		min = 0
		// fmt.Printf("%v %v return: %#v, ", len(indent), indent, min)
		return
	}

	// 表示不含有相同的字符
	min = math.MaxInt

	for i, sticker := range stickers {
		rest := minus(target, sticker, i, indent)
		if len(rest) != len(target) {
			m := process(rest, stickers, indent+"\t")
			// fmt.Printf("min: %v -> ", min)
			min = Min(min, m) // 如果剪除了字符，则表示是可以拼接的
			// fmt.Printf("%v\n", min)
		} else {
			// fmt.Printf("%v %v skip, min: %#v\n", len(indent), indent+"\t", min)
		}
	}

	// 如果可以拼接，证明找到了一张贴纸
	if min != math.MaxInt {
		min += 1
	}
	// fmt.Printf("%v %v return: %#v, ", len(indent), indent, min)
	return
}

func minus(target string, sticker string, index int, indent string) string {
	count := 26

	targetCounter := make([]int, count)
	for _, c := range target {
		targetCounter[c-'a'] += 1
	}

	stickerCounter := make([]int, count)
	for _, c := range sticker {
		stickerCounter[c-'a'] += 1
	}

	restLength := 0
	rest := make([]rune, count)
	for i := 0; i < count; i++ {
		if targetCounter[i] == 0 {
			continue
		}

		targetCounter[i] -= stickerCounter[i]
		for targetCounter[i] > 0 {
			rest[restLength] = 'a' + rune(i)
			targetCounter[i]--
			restLength++
		}
	}

	result := string(rest[:restLength])
	// fmt.Printf("%v %v - %#v = %#v\n", len(indent), indent, sticker, result)
	return result
}
