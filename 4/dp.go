package main

/*
 * 有一个由全是数字 (0123456789) 组成的字符串
 * 现要求将其转换成字母, 规则如下
 * 1 -> A
 * 2 -> B
 * 3 -> C
 * ...
 * 26 -> Z
 * 问有多少种转换结果
 *
 * 例如: 111 可以转换成 AAA, KA, AK, 所以有3种转换结果
 *
 * @digits 由多个数字组成的字符串
 * @return 有多少种方式可以将数字转换成字母
 */
func ConvertToLetterStringDp(digits string) int {
	if len(digits) == 0 {
		return 0
	}

	length := len(digits)

	// 递归只有一个变化参数，所以是一维
	dp := make([]int, length+1)

	// 递归终止条件
	dp[length] = 1

	// 看初始化和返回值，以及它依赖后面的两个位置，所以应该是从右往左计算
	for index := length - 1; index >= 0; index-- {
		if digits[index] == '0' {
			continue
		}

		// 单转
		dp[index] = dp[index+1]
		// 和下一个字符一起转
		if index+1 < length && (digits[index]-'0')*10+digits[index+1]-'0' < 27 {
			dp[index] += dp[index+2]
		}
	}

	// 递归首次调用时的参数
	return dp[0]
}
