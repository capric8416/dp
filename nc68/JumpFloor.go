package main

/*
NC68 跳台阶

知识点
递归
动态规划
记忆化搜索

描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

数据范围：1 <= n <= 40
要求：时间复杂度：O(n) ，空间复杂度： O(1)

示例1
输入：
2
返回值：
2

说明：
青蛙要跳上两级台阶有两种跳法，分别是：先跳一级，再跳一级或者直接跳两级。因此答案为2

示例2
输入：
7

返回值：
21
*/

/**
 *
 * @param number int整型，有 number 级台阶
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here

	// dp := make([]int, number+1)
	// dp[number] = 1
	// dp[number-1] = 1

	prevPrev := 1
	prev := 1
	current := 1

	for index := number - 2; index >= 0; index-- {
		// dp[index] = dp[index+1] + dp[index+2]
		current = prevPrev + prev
		prevPrev = prev
		prev = current
	}
	return current
	// return dp[0]

	// return process(0, number)
}

func process(current int, number int) int {
	// 变参只有current，变化范围是[0, number]，构建一维dp即可
	// 依赖current+1, current+2，说明只保留两个前置即可
	// 初始dp[number] = 0, dp[number-1] = 0
	// 从右往左计算 dp[number-2...0]

	// 如果跳到了最后一级，说明找到了一种跳法
	if current == number {
		return 1
	}
	// 如果越界，说明了此跳法无效
	if current > number {
		return 0
	}

	// 要么跳一级，要么跳两级，两种跳法之和
	return process(number, current+1) + process(number, current+2)
}

// func main() {
// 	fmt.Println(jumpFloor(2))
// 	fmt.Println(jumpFloor(7))
// }
