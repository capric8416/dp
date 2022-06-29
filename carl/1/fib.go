package main

import "fmt"

func fibDfs(n int) int {
	if n < 2 {
		return n
	}

	return fibDfs(n-1) + fibDfs(n-2)
}

func fibDp1(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibDp2(n int) int {
	if n < 2 {
		return n
	}

	dp0 := 0
	dp1 := 1
	for i := 2; i <= n; i++ {
		dp3 := dp1 + dp0
		dp0 = dp1
		dp1 = dp3
	}
	return dp1
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", fibDfs(i))
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", fibDp1(i))
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", fibDp2(i))
	}
	fmt.Println()
}
