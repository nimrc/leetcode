package solution

import . "github.com/nimrc/leetcode/pkg/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	curr := head
	var list *ListNode

	for curr != nil {
		list = &ListNode{
			Val:  curr.Val,
			Next: list,
		}
		curr = curr.Next
	}

	return list
}
