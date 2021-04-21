package solution

import (
	. "github.com/nimrc/leetcode/pkg/linkedlist"
)

// @link https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		}

		if pB == nil {
			pB = headA
		}

		pA = pA.Next
		pB = pB.Next

	}

	return pA
}
