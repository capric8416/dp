package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 表示把[0,i]物品放进重量为j的背包，获得的最大价值
	// dp[j] 表示限重为j的背包，最大可以背多少价值的物品 (总共有0到i号物品，每个物品只有一个)

	// 两个维度的01背包，i维表示i个0，j维表示j个1
	// dp[i][j] 表示i个0, j个1的背包，最大可以装进多少个s
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// dp[i][j] = max(dp[i-1][j], dp[i-1][j - weight[i]] + value[i])
	//    dp[j] = max(     dp[j], dp[i-1][j - weight[i]] + value[i])
	for _, s := range strs {
		// 统计0和1的个数
		var zeros, ones int
		for _, c := range s {
			if c == '0' {
				zeros++
			} else {
				ones++
			}
		}

		// 两层循环顺序可以交换
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- { // s可以放进去，所以子集加1
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("4 ==", findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println("2 ==", findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
