package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	half := sum / 2 // 努力分割成和相同的两个子集， 向下取整，所以sum-dp[half]一定大于等于dp[half]

	dp := make([]int, half+1)
	for i := 0; i < len(stones); i++ {
		for j := half; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[half] - dp[half]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("1 ==", lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("5 ==", lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}
