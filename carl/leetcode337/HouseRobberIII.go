package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	without, with := dfs(root)
	return max(without, with)
}

func dfs(root *TreeNode) (without int, with int) {
	if root == nil {
		return 0, 0
	}

	// 后序遍历
	withoutLeft, withLeft := dfs(root.Left)
	withoutRight, withRight := dfs(root.Right)

	// 偷根节点 + 不偷左节点 + 不偷右节点
	with = root.Val
	with += withoutLeft + withoutRight

	// 不偷根节点 = 偷左节点最大值 + 偷右节点最大值
	without = max(withoutLeft, withLeft) + max(withoutRight, withRight)

	return without, with
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
