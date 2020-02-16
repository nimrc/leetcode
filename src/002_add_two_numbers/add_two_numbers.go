package solution

import (
	. "github.com/fyibmsd/leetcode/common"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list, tail *ListNode
	tmp := 0

	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		if tail == nil {
			list = &ListNode{Val: tmp % 10, Next: nil}
			tail = list
		} else {
			tail.Next = &ListNode{Val: tmp % 10, Next: nil}
			tail = tail.Next
		}

		tmp /= 10
	}

	return list
}
