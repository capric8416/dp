package main

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dp := make([]int, size) // dp[i]，以i结尾的连续子数组的最大和
	dp[0] = nums[0]

	sum := dp[0]
	for i := 1; i < size; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > sum {
			sum = dp[i]
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
