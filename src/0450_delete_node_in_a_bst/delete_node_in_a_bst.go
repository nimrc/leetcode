package solution

import . "github.com/nimrc/leetcode/pkg/binarytree"

func deleteNode(root *TreeNode, key int) *TreeNode {
	prev := root
	curr := prev

	for curr.Val != key {
		prev = curr
		if curr.Val < key {
			curr = curr.Right
		} else {
			curr = curr.Left
		}

		if curr == nil {
			return root
		}
	}

	if prev.Left == curr {
		prev.Left = curr.Right

	} else {
		prev.Right = curr.Left
	}

	return curr
}
