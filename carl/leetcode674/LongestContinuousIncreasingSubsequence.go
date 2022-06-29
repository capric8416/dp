package main

func findLengthOfLCIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	// dp[i]表示以i结尾的连续递增子序列的最大长度
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	length := 0
	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > length {
				length = dp[i]
			}
		}
	}

	return length
}

func main() {

}
