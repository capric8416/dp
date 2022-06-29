package main

import "container/heap"

/*
 *
 * 给定一个数据 washMachineTime, washMachineTime[i]代表第i号咖啡泡一杯咖啡的时间
 * 给定一个正数 customers, 表示 customers 个人等待着咖啡机泡咖啡, 每台咖啡机只能泡咖啡
 * 只有一台洗咖啡杯机, 一次只能洗一个杯子, 时间耗费 washTime, 洗完才能洗下一杯
 * 每个咖啡杯也可以自己挥发干净, 时间耗费 dryTime, 咖啡杯可以并行挥发
 * 假设所有人拿到咖啡之后立刻喝干净, 时间耗费 0
 * 返回从开始等待泡咖啡所有咖啡杯变干净的最短时间
 *
 */

func CoffeeDp(washMachineTime []int, customers int, washTime int, dryTime int) int {
	// 先针对咖啡机进行排队

	h := make(TimeHeap, len(washMachineTime))
	for i, v := range washMachineTime {
		h[i] = Time{timePoint: 0, workTime: v}
	}
	heap.Init(&h)

	// 喝完的时间点 (即可以开始洗咖啡杯或者让其挥发了)
	drinks := make([]int, customers)
	for i := 0; i < customers; i++ {
		v := heap.Pop(&h).(Time)
		drinks[i] = v.timePoint + v.workTime
		heap.Push(&h, Time{timePoint: v.timePoint + v.workTime, workTime: v.workTime})
	}

	// 考虑最杯情况, 所有的杯子都用来洗
	allWashTime := 0
	for _, d := range drinks {
		allWashTime = Max(allWashTime, d) + washTime
	}

	size := len(drinks)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, allWashTime+1)
	}

	for washLine, index := 0, size-1; washLine <= allWashTime; washLine++ {
		// 洗咖啡杯机开始可用的时间, 与喝完的时间点取大值, 再加上洗的耗时
		t1 := Max(washLine, drinks[index]) + washTime
		// 喝完的时间点, 加上挥发的时间
		t2 := drinks[index] + dryTime
		// 取小值
		dp[index][washLine] = Min(t1, t2)
	}

	for index := size - 2; index >= 0; index-- {
		for washLine := 0; washLine <= allWashTime; washLine++ {
			// 洗
			// 当前这杯
			wash := Max(washLine, drinks[index]) + washTime
			if wash > allWashTime {
				continue
			}
			// index...杯
			next1 := dp[index+1][wash]

			// 挥发
			// 当前这杯
			dry := drinks[index] + dryTime
			// index...杯
			next2 := dp[index+1][washLine]

			dp[index][washLine] = Min(Max(wash, next1), Max(dry, next2))
		}
	}

	return dp[0][0]
}
