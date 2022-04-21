package main

/*
 * 有一堆物品，它们的每一个的重量分别是 weights[0...]
 * 它们的每一个的价值分别是 values[0...]
 * 有一个背包，它最最多可以装 bagCapacity 重的物品
 * 问：最大可以装多少价值的物品？
 *
 * @weights 物品的重量数组
 * @values 物品的价值数组
 * @bagCapacity 背包的容量（最大载重）
 * @return 背包最大可以装多少价值的物品？
 */
func KnapsackDp(weights []int, values []int, bagCapacity int) int {
	if len(weights) != len(values) || bagCapacity < 0 || weights == nil || values == nil {
		return 0
	}

	size := len(weights)

	// 递归有两个可变参数，所以是二维
	// index 变化范围是 [0, size]
	// leftBagCapacity 变化范围是 [0, leftBagCapacity]
	dp := make([][]int, size+1)
	for index := 0; index <= size; index++ {
		dp[index] = make([]int, bagCapacity+1)
	}

	// 递归终止条件为 leftBagCapacity <= 0 || index == size
	// 所以dp[size][...] = 0，默认初始化都为0，无需初始化
	// 最后一排都为0, 左侧列越界都无效，目录在第0行，第 bagCapacity 列
	// 所以从下到上，从左到右计算

	for index := size - 1; index >= 0; index-- {
		for leftBagCapacity := 0; leftBagCapacity <= bagCapacity; leftBagCapacity++ {
			v1 := dp[index+1][leftBagCapacity]
			v2 := 0
			if leftBagCapacity-weights[index] >= 0 {
				v2 = values[index] + dp[index+1][leftBagCapacity-weights[index]]
			}
			dp[index][leftBagCapacity] = Max(v1, v2)
		}
	}

	return dp[0][bagCapacity]
}
