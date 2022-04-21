package main

/*
 * @description
 * 给定一个整形数组 arr，代表数值不同的纸牌排成一条线
 * 玩家A和玩家B依次拿走每一张牌
 * 规定玩家A先拿，玩家B后拿
 * 但是每个玩家每次只能拿走最左或最右的纸牌
 * 玩家A和玩家B都绝顶聪明
 * 请返回最后获胜者的分数
 *
 * @arr: 排成一条线的数值不同的纸牌
 * @left: 左边界（可达）
 * @right: 右边界 （可达）
 * @return: 返回最后获胜者的分数
 */
func CardsInLineDp(arr []int, left, right int) int {
	if len(arr) == 0 {
		return 0
	}

	// 观察递归函数，有两个变化参数，所以dp是二维的
	// left 变化范围是 [0, size-1]
	// right 变化范围是 [0, size-1]
	// left <= right，所以对角线以下的空间用不到
	size := len(arr)
	firstDp := make([][]int, size)
	secondDp := make([][]int, size)
	for l := 0; l < size; l++ {
		firstDp[l] = make([]int, size)
		secondDp[l] = make([]int, size)

		// 观察递归终止条件
		// 先手dp，对角线初始化为arr[l]
		firstDp[l][l] = arr[l]

		// 后手dp，对角线默认就是0, 无需初始化
	}

	// // 先手dp依赖后手dp左侧和下方
	// // 后手dp依赖先手dp的左侧和下方
	// // 初始化对角线右侧
	// for l := 0; l < size-1; l++ {
	// 	r := l + 1
	// 	secondDp[l][r] = Min(firstDp[l+1][r], firstDp[l][r-1])
	// 	firstDp[l][r] = Max(arr[l]+secondDp[l+1][r], arr[r]+secondDp[l][r-1])
	// }

	// 从下到上，从左至右，跳过对角线1格
	for l := size - 1; l >= 0; l-- {
		for r := l + 1; r < size; r++ {
			firstDp[l][r] = Max(arr[l]+secondDp[l+1][r], arr[r]+secondDp[l][r-1])
			secondDp[l][r] = Min(firstDp[l+1][r], firstDp[l][r-1])
		}
	}

	return Max(firstDp[0][size-1], secondDp[0][size-1])
}
