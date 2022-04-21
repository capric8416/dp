package main

func KnapsackRecurse(weights []int, values []int, bagCapacity int) int {
	if len(weights) != len(values) || bagCapacity < 0 || weights == nil || values == nil {
		return 0
	}

	return process(0, bagCapacity, weights, values)
}

// 背包还剩下 leftBagCapacity 容量，针对 index 号物品，拿或者不拿，最大可以取得多少价值
func process(index int, leftBagCapacity int, weight []int, values []int) int {
	// 背包已经装不上东西了
	if leftBagCapacity <= 0 || index == len(weight) {
		return 0
	}

	// 不拿 index 号物品，可以取得的最大价值
	v1 := process(index+1, leftBagCapacity, weight, values)
	// 拿 index 号物品，可以取得的最大价值
	v2 := 0
	if leftBagCapacity-weight[index] >= 0 { // 只有当背包空闲容量还可以装的得下 index 号物品时，才考虑装它
		v2 = values[index] + process(index+1, leftBagCapacity-weight[index], weight, values)
	}

	return Max(v1, v2)
}
