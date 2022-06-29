package main

import (
	"container/heap"
)

type Time struct {
	timePoint int
	workTime  int
}

type TimeHeap []Time

func (h TimeHeap) Len() int {
	return len(h)
}

func (h TimeHeap) Less(i, j int) bool {
	return h[i].timePoint+h[i].workTime < h[j].timePoint+h[j].workTime
}

func (h TimeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TimeHeap) Push(x interface{}) {
	*h = append(*h, x.(Time))
}

func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CoffeeRecurse(washMachineTime []int, customers int, washTime int, dryTime int) int {
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

	return process(0, 0, drinks, washTime, dryTime)
}

/*
 * @index drinks 喝完的时间点数据索引
 * @washLine 洗咖啡机开始可用的时间
 * @drinks 喝完的时间点
 * @washTime 洗咖啡杯需要的时间
 * @dryTime 挥发咖啡杯需要的时间
 * @return 所以的咖啡杯变干净的最少时间
 */
func process(index int, washLine int, drinks []int, washTime int, dryTime int) int {
	// 最后一个杯子了
	if index == len(drinks)-1 {
		// 洗咖啡杯机开始可用的时间, 与喝完的时间点取大值, 再加上洗的耗时
		t1 := Max(washLine, drinks[index]) + washTime
		// 喝完的时间点, 加上挥发的时间
		t2 := drinks[index] + dryTime
		// 取小值
		return Min(t1, t2)
	}

	// 洗
	// 当前这杯
	wash := Max(washLine, drinks[index]) + washTime
	// index...杯
	next1 := process(index+1, wash, drinks, washTime, dryTime)

	// 挥发
	// 当前这杯
	dry := drinks[index] + dryTime
	// index...杯
	next2 := process(index+1, washLine, drinks, washTime, dryTime)

	return Min(Max(wash, next1), Max(dry, next2))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
