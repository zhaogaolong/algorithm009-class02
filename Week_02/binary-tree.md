# 二叉树遍历实现总结

## 前序遍历

https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil{
        return nil
    }
    res := []int{}
    stack := []*TreeNode{root}
    for len(stack) > 0{
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil{
            res = append(res, node.Val)
            stack = append(stack, node.Right)
            stack = append(stack, node.Left)
        }
    }
    return res

}
```

多叉树遍历

https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

```go
// 多叉树
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil{
        return nil
    }
    res := []int{}
    stack := []*Node{root}
    for len(stack) > 0{
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil{
            res = append(res, node.Val)
            for i := len(node.Children)-1; i >= 0; i--{
                stack = append(stack, node.Children[i])
            }
        }
    }
    return res

}

```


## 中序遍历

https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil{
        return nil
    }
    res := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0{
        for root != nil{
            stack = append(stack, root)
            root = root.Left
        }
        // 弹栈
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // 放数据
        res = append(res, root.Val)
        root = root.Right
    }
    return res

}

```

## 后续遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil{
        return nil
    }

    res := []int{}
    stack := []*TreeNode{root}
    for len(stack) > 0{
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        res = append([]int{node.Val}, res...)
        if node.Left != nil{
            stack = append(stack, node.Left)
        }
        if node.Right != nil{
            stack = append(stack, node.Right)
        }
    }
    return res
}

// N叉树后序遍历

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil{
        return nil
    }
    res := []int{}
    stack := []*Node{root}
    for len(stack) > 0{
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        res = append([]int{node.Val}, res...)
        for i := range node.Children{
            stack = append(stack, node.Children[i])
        }
    }
    return res
}
```


// 二叉树的层序遍历


// N 叉树层序遍历, 迭代实现
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

```go
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

```
