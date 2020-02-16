package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Value() int {
	if list.Next == nil {
		return list.Val
	}

	return list.Next.Value()*10 + list.Val
}

func NewList(elem ...int) *ListNode {
	var node *ListNode

	for _, val := range elem {
		if node == nil {
			node = &ListNode{Val: val, Next: nil}
		} else {
			node = &ListNode{Val: val, Next: node}
		}
	}

	return node
}
