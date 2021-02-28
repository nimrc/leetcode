package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	visited := make(map[*TreeNode]bool)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if (item.Left != nil || item.Right != nil) && !visited[item] {
			visited[item] = true

			stack = append(stack, item)

			if item.Right != nil {
				stack = append(stack, item.Right)
			}

			if item.Left != nil {
				stack = append(stack, item.Left)
			}
		} else {
			res = append(res, item.Val)
		}
	}

	return res
}
