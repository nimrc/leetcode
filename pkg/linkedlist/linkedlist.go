package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	return &ListNode{
		Val:  values[0],
		Next: NewList(values[1:]),
	}
}

func (list *ListNode) Append(head *ListNode) *ListNode {
	curr := list

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = head

	return list
}

func (list *ListNode) Dump() {
	for curr := list; curr != nil; curr = curr.Next {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Printf(" -> ")
		}
	}
}
