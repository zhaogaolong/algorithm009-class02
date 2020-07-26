package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left == nil {
		return r + 1
	} else if root.Right == nil {
		return l + 1
	} else {
		return min(l, r) + 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
