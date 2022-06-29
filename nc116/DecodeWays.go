package main

import "fmt"

/*
 * NC116 把数字翻译成字符串
 *
 * 描述
 * 有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
 * 我们把一个字符串编码成一串数字，再考虑逆向编译成字符串。
 * 由于没有分隔符，数字编码成字母可能有多种编译结果，例如 11 既可以看做是两个 'a' 也可以看做是一个 'k' 。但 10 只可能是 'j' ，因为 0 不能编译成任何结果。
 * 现在给一串数字，返回有多少种可能的译码结果
 *
 * 数据范围：字符串长度满足 0<n≤90
 * 进阶：空间复杂度 O(n)，时间复杂度 O(n)
 * 示例1
 * 输入：
 * "12"
 * 返回值：
 * 2
 * 说明：
 * 2种可能的译码结果（”ab” 或”l”）
 *
 * 示例2
 * 输入：
 * "31717126241541717"
 * 返回值：
 * 192
 * 说明：
 * 192种可能的译码结果
 */

/**
 * 解码
 * @param nums string字符串 数字串
 * @return int整型
 */
func solve(nums string) int {
	// write code here
	// return process(nums, 0)

	size := len(nums)
	dp := make([]int, size+1)
	dp[size] = 1
	for i := size - 1; i >= 0; i-- {
		if nums[i] == '0' {
			dp[i] = 0
			continue
		}
		dp[i] = dp[i+1]
		if i+1 < size && (nums[i]-'0')*10+(nums[i+1]-'0') < 27 {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}

func process(nums string, index int) int {
	if index == len(nums) { // 空字符串可以翻译成空字符串，只有一种转换方式
		return 1
	}

	if nums[index] == '0' { // 0字符，或者以0开始的两个字符都转换不了
		return 0
	}

	count := process(nums, index+1)
	if index+1 < len(nums) && (nums[index]-'0')*10+(nums[index+1]-'0') < 27 {
		count += process(nums, index+2)
	}

	return count
}

func main() {
	fmt.Println(solve("10"), 1)
	fmt.Println(solve("12"), 2)
	fmt.Println(solve("31717126241541717"), 192)
}
