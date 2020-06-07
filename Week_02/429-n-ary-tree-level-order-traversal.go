package week02

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	level := 0
	queue := []*Node{root}

	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] == nil {
				continue
			}
			res[level] = append(res[level], queue[i].Val)
			for _, n := range queue[i].Children {
				queue = append(queue, n)
			}
		}
		queue = queue[counter:]
		level++
	}
	return res
}
