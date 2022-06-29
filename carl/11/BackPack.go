package main

import "fmt"

func BackPack(bagCapacity int, itemWeights []int, itemValues []int) int {
	itemCount := len(itemWeights)

	dp := make([][]int, itemCount)
	for i := 0; i < itemCount; i++ {
		dp[i] = make([]int, bagCapacity+1)
	}

	// 初始化dp[i][0]为0(创建dp数据时已完成), 因为背包容量为0, 装不下任何任何物品
	// 初始化dp[0][j], 第0行, 如果背包放的下物品i, 则背包最大价值为物品i的价值
	for j := 1; j <= bagCapacity; j++ {
		if j >= itemWeights[0] {
			dp[0][j] = itemValues[0]
		}
	}

	// 循环可以调换顺序, 因为依赖的是左上角
	for i := 1; i < itemCount; i++ { // 遍历物品
		for j := 1; j <= bagCapacity; j++ { // 遍历背包容量
			if j < itemWeights[i] { // 背包装不下物品i
				dp[i][j] = dp[i-1][j] // 正上方
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-itemWeights[i]]+itemValues[i]) // 左上角
			}
		}
	}

	return dp[itemCount-1][bagCapacity]
}

func BackPack1(bagCapacity int, itemWeights []int, itemValues []int) int {
	itemCount := len(itemWeights)

	dp := make([]int, bagCapacity+1)

	for i := 0; i < itemCount; i++ { // 遍历物品
		for j := bagCapacity; j >= itemWeights[i]; j-- { // 遍历背包容量
			dp[j] = max(dp[j], dp[j-itemWeights[i]]+itemValues[i])
		}
	}

	// // 如果交互内外层循环，则背包只会放一个物品
	// for j := bagCapacity; j > 0; j-- { // 遍历背包容量
	// 	for i := 0; i < itemCount; i++ { // 遍历物品
	// 		if j >= itemWeights[i] {
	// 			dp[j] = max(dp[j], dp[j-itemWeights[i]]+itemValues[i])
	// 			continue
	// 		}
	// 		break
	// 	}
	// }

	// // 如果内层循环从小到大遍历背包重量，则物品会被多次放入
	// for i := 0; i < itemCount; i++ { // 遍历物品
	// 	for j := itemWeights[i]; j <= bagCapacity; j++ { // 遍历背包容量
	// 		dp[j] = max(dp[j], dp[j-itemWeights[i]]+itemValues[i])
	// 	}
	// }

	return dp[bagCapacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("35 ==", BackPack1(4, []int{1, 3, 4}, []int{15, 20, 30}))
	fmt.Println("10 ==", BackPack1(6, []int{1, 2, 3, 4, 5}, []int{2, 4, 4, 5, 6}))
}
