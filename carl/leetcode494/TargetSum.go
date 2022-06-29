package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	// 把数组分割成两个子集，分别为subtrahend(所有添加正号的数组成减数)和minuend（所有添加负号的数组成被减数），则
	// subtrahend - minuend = target
	// subtrahend + minuend = sum
	// 两式相加除2可得subtrahend = (target+sum)/2
	// 即转化为了01背包， 背包容量为(target+sum)/2

	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 目标值大于总和
	// 不可均分target+sum
	if abs(target) > sum || (target+sum)%2 != 0 {
		return 0
	}

	bagWeight := (target + sum) / 2

	dp := make([]int, bagWeight+1) // dp[j]表示装满重量为j的背包，总共有多少种方法
	dp[0] = 1                      // 背包重量为0，只有一种方法，就是什么都不装，这里很重要，要不然dp数组计算过程中就全是0
	for i := 0; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]] // 组合
		}
	}

	return dp[bagWeight]
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	fmt.Println("1 ==", findTargetSumWays([]int{1}, 1))
	fmt.Println("5 ==", findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println("1 ==", findTargetSumWays([]int{1}, 1))
	fmt.Println("0 ==", findTargetSumWays([]int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}, 0))
}
