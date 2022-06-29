package main

import (
	"fmt"
	"math"
)

func FindGreatestSumOfSubArray1(array []int) int {
	return process(array)
}

func process(array []int) int {
	if len(array) < 0 {
		return 0
	}

	maxSum := math.MinInt
	for i := 0; i < len(array); i++ {
		for step := 0; i+step < len(array); step++ {
			sum := 0
			for j := i; j <= i+step; j++ {
				sum += array[j]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	fmt.Println(FindGreatestSumOfSubArray1([]int{-2, -8, -1, -5, -9}))
	fmt.Println(FindGreatestSumOfSubArray1([]int{1, -2, 3, 10, -4, 7, 2, -5}), 18)
	fmt.Println(FindGreatestSumOfSubArray1([]int{2}), 2)
	fmt.Println(FindGreatestSumOfSubArray1([]int{-10}), -10)
}
