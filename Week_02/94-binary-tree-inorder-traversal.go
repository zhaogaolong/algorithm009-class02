package week02

// 递归实现

var res []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res = make([]int, 0)
	dfs94(root)
	return res
}

func dfs94(root *TreeNode) {
	if root == nil {
		return
	}
	dfs94(root.Left)
	res = append(res, root.Val)
	dfs94(root.Right)
}

// 栈实现
func inorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		i := len(stack) - 1
		res = append(res, stack[i].Val)
		root = stack[i]
		stack = stack[:i]
	}
	return res

}
