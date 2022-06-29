package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func LargeDFS(root *TreeNode, k *int, target *int) {
	if root == nil {
		return
	}

	LargeDFS(root.Right, k, target)

	if (*k)--; *k == 0 {
		*target = root.Value
		return
	}

	LargeDFS(root.Left, k, target)
}

func SamllDFS(root *TreeNode, k *int, target *int) {
	if root == nil {
		return
	}

	SamllDFS(root.Left, k, target)

	if (*k)--; *k == 0 {
		*target = root.Value
		return
	}

	SamllDFS(root.Right, k, target)
}

func FindKthLargeNumber(root *TreeNode, k int) int {
	// 二叉搜索树
	// 从大到小遍历：右 -> 中 -> 左
	target := -1
	LargeDFS(root, &k, &target)
	return target
}

func FindKthSmallNumber(root *TreeNode, k int) int {
	// 二叉搜索树
	// 从大到小遍历：左 -> 中 -> 右
	target := -1
	SamllDFS(root, &k, &target)
	return target
}

func main() {
	one := TreeNode{Value: 1}
	two := TreeNode{Value: 2, Left: &one}
	four := TreeNode{Value: 4}
	three := TreeNode{Value: 3, Left: &two, Right: &four}

	nine := TreeNode{Value: 9}
	eight := TreeNode{Value: 8, Right: &nine}
	six := TreeNode{Value: 6}
	seven := TreeNode{Value: 7, Left: &six, Right: &eight}

	five := TreeNode{Value: 5, Left: &three, Right: &seven}

	for i := 1; i < 20; i++ {
		fmt.Printf("%v ", FindKthLargeNumber(&five, i))
	}
	fmt.Println()
	for i := 1; i < 20; i++ {
		fmt.Printf("%v ", FindKthSmallNumber(&five, i))
	}
	fmt.Println()

	four1 := TreeNode{Value: 4}
	six1 := TreeNode{Value: 6}
	five1 := TreeNode{Value: 5, Left: &four1, Right: &six1}
	two1 := TreeNode{Value: 2}
	three1 := TreeNode{Value: 3, Left: &two1, Right: &five1}
	eight1 := TreeNode{Value: 8}
	seven1 := TreeNode{Value: 7, Left: &three1, Right: &eight1}
	one1 := TreeNode{Value: 1, Right: &seven1}

	for i := 1; i < 20; i++ {
		fmt.Printf("%v ", FindKthLargeNumber(&one1, i))
	}
	fmt.Println()
	for i := 1; i < 20; i++ {
		fmt.Printf("%v ", FindKthSmallNumber(&one1, i))
	}
	fmt.Println()
}
