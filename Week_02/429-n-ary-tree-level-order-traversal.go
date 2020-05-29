package week02

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*Node{root}
	level := 0
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})

		for i := 0; i < counter; i++ {
			n := queue[i]
			if n == nil {
				continue
			}
			res[level] = append(res[level], n.Val)
			for _, n := range queue[i].Children {
				queue = append(queue, n)
			}
		}
		queue = queue[counter:]
		level++
	}

	return res
}
