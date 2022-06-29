package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * @param prices: stock prices of day 0...i
 * @return: return the maximum profit
 */
func GetAns(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	h := make(IntHeap, 0)
	heap.Init(&h)

	profit := 0
	for _, price := range prices {
		heap.Push(&h, price)
		if len(h) > 0 && price >= h[0] {
			profit += price - h[0]
			heap.Pop(&h)
			heap.Push(&h, price)
		}
	}

	return profit
}

func main() {
	fmt.Println("16 ==", GetAns([]int{1, 2, 10, 9}))
	fmt.Println("5 ==", GetAns([]int{9, 5, 9, 10, 5}))
}
