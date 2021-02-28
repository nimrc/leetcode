package solution

import . "github.com/nimrc/leetcode/pkg/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list, curr *ListNode
	var tmp int

	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		if curr == nil {
			list = &ListNode{Val: tmp % 10}
			curr = list
		} else {
			curr.Next = &ListNode{Val: tmp % 10}
			curr = curr.Next
		}
		tmp /= 10
	}

	return list
}
