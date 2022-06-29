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
 * 要求：时间复杂度 O(nlogn)， 空间复杂度 O(n)
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

	if len(arr) == 0 {
		return 0
	}

	results := make([]int, 1) // 存储最长上升子序列
	results[0] = arr[0]       // 默认添加原数组首个元素

	for i := 1; i < len(arr); i++ {
		if arr[i] > results[len(results)-1] { // 如果大于最长上升子序列的最大的元素，则直接将其追加到末尾
			results = append(results, arr[i])
			continue
		}

		left, right := 0, len(results)-1
		for left <= right { // 否则二分查找，插入到 lower bound 位置
			mid := left + (right-left)/2
			if results[mid] > arr[i] {
				right = mid - 1
			} else if results[mid] < arr[i] {
				left = mid + 1
			} else {
				left = mid
				break
			}
		}
		results[left] = arr[i]
	}

	return len(results)
}

func main() {
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(LIS([]int{6, 3, 1, 5, 2, 3, 7}))
	fmt.Println(LIS([]int{10, 9, 1, 5, 7, 2, 3, 4}))
}
