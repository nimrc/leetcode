package binarytree

import (
	"github.com/davecgh/go-spew/spew"
	"math"
)

const (
	None = math.MaxInt64
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode) {
	curr := root
	q := make([]*TreeNode, 0)

	for curr != nil {
		q = append(q, curr)

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}

	}

	spew.Dump(q)
}

func BuildTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	return build(values, 0)
}

func build(values []int, idx int) *TreeNode {
	if idx >= len(values) || values[idx] == None {
		return nil
	}

	return &TreeNode{
		Val:   values[idx],
		Left:  build(values, 2*idx+1),
		Right: build(values, 2*idx+2),
	}
}
