package main

import "fmt"

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	if len(array) < 0 {
		return 0
	}

	count := len(array)
	dp := make([]int, count)
	dp[0] = array[0]

	maxSum := dp[0]
	for i := 1; i < count; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + array[i]
		} else {
			dp[i] = array[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum
}

func main() {
	fmt.Println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
}
