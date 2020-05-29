package week02

type Nodes struct {
	Val      int
	Children []*Nodes
}

func preorder(root *Nodes) []int {
	res := []int{}
	stack := []*Nodes{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}

		}
	}
	return res
}
