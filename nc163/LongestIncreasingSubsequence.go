package main

import "fmt"

/*
 * NC163 最长上升子序列(一)
 *
 * 描述
 * 给定一个长度为 n 的数组 arr，求它的最长严格上升子序列的长度。
 * 所谓子序列，指一个数组删掉一些数（也可以不删）之后，形成的新数组。例如 [1,5,3,7,3] 数组，其子序列有：[1,3,3]、[7] 等。但 [1,6]、[1,3,5] 则不是它的子序列。
 * 我们定义一个序列是 严格上升 的，当且仅当该序列不存在两个下标 i 和 i 满足 i<j 且 arr[i] >= arr[j]
 *
 * 数据范围： 0≤n≤1000
 * 要求：时间复杂度 O(n^2)， 空间复杂度 O(n)
 * 示例1
 * 输入：
 * [6,3,1,5,2,3,7]
 *
 * 返回值：
 * 4
 *
 * 说明：
 * 该数组最长上升子序列为 [1,2,3,7] ，长度为4
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 给定数组的最长严格上升子序列的长度。
 * @param arr int整型一维数组 给定的数组
 * @return int整型
 */
func LIS(arr []int) int {
	// write code here
	// return process(arr, -1, 0)

	size := len(arr)
	dp := make([]int, size) // dp[i] 表示以第i个数组元素结尾的部分，最大上升子序列长度

	maxLength := 0
	for currIndex := 0; currIndex < size; currIndex++ { // 从左到右计算 dp[0], dp[1], ..., dp[size-1]
		dp[currIndex] = 1                                        // 自身一个元素可构成长度为一的上长子序列
		for prevIndex := 0; prevIndex < currIndex; prevIndex++ { // 查找当前位置之前的区间
			if arr[currIndex] > arr[prevIndex] { // 如果当前位置元素大于之前位置的元素，长度可能会加1
				if dp[prevIndex]+1 > dp[currIndex] { // dp[0], ..., dp[prevIndex] 都计算过了
					dp[currIndex] = dp[prevIndex] + 1
				}
			}
		}
		if dp[currIndex] > maxLength {
			maxLength = dp[currIndex]
		}
	}

	return maxLength
}

func process(arr []int, prevIndex int, currIndex int) int {
	if currIndex == len(arr) { // 越界，找不到
		return 0
	}

	count := 0
	if prevIndex < 0 || arr[prevIndex] < arr[currIndex] { // 前面没有数字；前面的数字小于当前数字，说明找到了一个递增
		count = 1 + process(arr, currIndex, currIndex+1)
	}

	c := process(arr, prevIndex, currIndex+1) // 前面的数字不小于当前数字，在后面的去找递增
	if c > count {
		count = c
	}

	return count
}

func main() {
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(LIS([]int{6, 3, 1, 5, 2, 3, 7}))
}
