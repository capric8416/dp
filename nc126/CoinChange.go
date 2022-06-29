package main

import (
	"fmt"
	"math"
)

/*
* NC126 兑换零钱(一)
*
* 描述
* 给定数组arr，arr中所有的值都为正整数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
* 如果无解，请返回-1.
*
* 数据范围：
* 数组大小满足 0 <= n <= 100000 ，
* 数组中每个数字都满足 0 < val <= 10000，
* 0 <= aim <= 5000
*
* 要求：时间复杂度 O(n x aim) ，空间复杂度 O(aim)。
*
* 示例1
* 输入：
* [5,2,3],20
* 返回值：
* 4
*
*
* 示例2
* 输入：
* [5,2,3],0
* 返回值：
* 0
*
*
* 示例3
* 输入：
* [3,5],2
* 返回值：
* -1
 */

/**
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
func minMoney(arr []int, aim int) int {
	// write code here

	// count := process(arr, aim)
	// if count == math.MaxInt32 {
	// 	count = -1
	// }
	// return count

	dp := make([]int, aim+1)
	for i := 1; i <= aim; i++ {
		count := math.MaxInt32
		for _, v := range arr {
			if i-v >= 0 && dp[i-v] != math.MaxInt32 {
				c := 1 + dp[i-v]
				if c < count {
					count = c
				}
			}
		}
		dp[i] = count
	}
	if dp[aim] == math.MaxInt32 {
		dp[aim] = -1
	}

	return dp[aim]
}

func process(arr []int, aim int) int {
	if aim == 0 {
		return 0 // 目标值为零，则最少需要0个货币数
	}

	count := math.MaxInt32 // 初始值为无数个，表示这条树的路径不通
	for _, v := range arr {
		leftAim := aim - v
		if leftAim >= 0 { // 如果减去当前值还有剩余，则递归
			c := process(arr, leftAim)
			if c < count {
				count = c
			}
		}
	}
	if count != math.MaxInt32 { // 这条递归树的路径走通了，深度(即货币数)+1
		count += 1
	}

	return count
}

func main() {
	fmt.Println(minMoney([]int{5, 2, 3}, 20), 4)
	fmt.Println(minMoney([]int{5, 2, 3}, 0), 0)
	fmt.Println(minMoney([]int{3, 5}, 2), -1)
}
