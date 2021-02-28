package solution

import (
	"github.com/davecgh/go-spew/spew"
	. "github.com/nimrc/leetcode/pkg/binarytree"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	node := deleteNode(root, 3)

	spew.Dump(node)
}
