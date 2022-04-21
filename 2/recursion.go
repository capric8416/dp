package main

func CardsInLineRecurse(arr []int, left, right int) int {
	if len(arr) == 0 {
		return 0
	}

	// 在全部的牌中执先手
	first := First(arr, 0, len(arr)-1)
	// 在全部的牌中执后手
	after := Second(arr, 0, len(arr)-1)
	// 返回获胜者的得分
	return Max(first, after)
}

func First(arr []int, left, right int) int {
	// 只剩下最后一张牌，先手直接拿走
	if left == right {
		return arr[left]
	}

	// 先拿左边的，下一步在剩下的牌中，处于后手
	l := arr[left] + Second(arr, left+1, right)
	// 先拿右边的，下一步在剩下的牌中，处于后手
	r := arr[right] + Second(arr, left, right-1)
	// 先手拿最大的牌
	return Max(l, r)
}

func Second(arr []int, left, right int) int {
	// 只剩下一张牌，肯定会被先手拿走
	if left == right {
		return 0
	}

	// 如果对方拿走了左边的牌
	l := First(arr, left+1, right)
	// 如果对方拿走了右边的牌
	r := First(arr, left, right-1)
	// 对方先手，不会让我拿最大的牌
	return Min(l, r)
}
