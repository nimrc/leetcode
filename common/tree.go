package common

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	None = math.MaxInt64
)

func Build(tree []int) *TreeNode {
	return build(tree, 0)
}

func build(tree []int, idx int) *TreeNode {
	if len(tree) < idx+1 {
		return nil
	}

	if tree[idx] == None {
		return nil
	}

	return &TreeNode{
		Val:   tree[idx],
		Left:  build(tree, 2*idx+1),
		Right: build(tree, 2*idx+2),
	}
}
