package main

import "fmt"

// n为正整数, 1...n
// 初始是站在0位置
func climbStairsDfs(n int) int {
	if n < 2 {
		return n
	}
	return climbStairsDfs(n-1) + climbStairsDfs(n-2)
}

func climbStairsDp1(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = 2
	for i := n - 2; i > 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
	}
	return dp[1]
}

func climbStairsDp2(n int) int {
	if n < 2 {
		return n
	}

	// dp := make([]int, n+1)
	dp1 := 1 // dp[1] = 1
	dp2 := 2 // dp[2] = 2
	for i := 3; i <= n; i++ {
		// dp[i] = dp[i-1] + dp[i-2] // 从0爬到i阶台阶, 总共有多少种方法; 最后一步爬到了i阶, 那上一步要么是从i-1, 要么是从i-2爬上来的
		dp3 := dp1 + dp2
		dp1 = dp2
		dp2 = dp3
	}
	return dp2 // return dp[n]
}

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%v ", climbStairsDfs(i))
	// }
	// fmt.Println()

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%v ", climbStairsDp1(i))
	// }
	// fmt.Println()

	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d ", climbStairsDp2(i))
	}
	fmt.Println()
}
