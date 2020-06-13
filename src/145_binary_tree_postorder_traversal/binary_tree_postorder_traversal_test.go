package solution

import (
	"fmt"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {

	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	res := postorderTraversal(tree)

	fmt.Println(res)
}
