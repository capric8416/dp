package main

import "fmt"

/**
 * @param m: An integer m denotes the size of a backpack
 * @param a: Given n items with size A[i]
 * @return: The maximum size
 */
func BackPack(m int, a []int) int {
	// write your code here

	dp := make([]int, m+1)
	for i := 0; i < len(a); i++ {
		for j := m; j >= a[i]; j-- {
			dp[j] = max(dp[j], dp[j-a[i]]+a[i])
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("9 ==", BackPack(10, []int{3, 4, 8, 5}))
	fmt.Println("12 ==", BackPack(12, []int{2, 3, 5, 7}))
	fmt.Println("5000 ==", BackPack(5000, []int{828, 125, 740, 724, 983, 321, 773, 678, 841, 842, 875, 377, 674, 144, 340, 467, 625, 916, 463, 922, 255, 662, 692, 123, 778, 766, 254, 559, 480, 483, 904, 60, 305, 966, 872, 935, 626, 691, 832, 998, 508, 657, 215, 162, 858, 179, 869, 674, 452, 158, 520, 138, 847, 452, 764, 995, 600, 568, 92, 496, 533, 404, 186, 345, 304, 420, 181, 73, 547, 281, 374, 376, 454, 438, 553, 929, 140, 298, 451, 674, 91, 531, 685, 862, 446, 262, 477, 573, 627, 624, 814, 103, 294, 388}))
}
