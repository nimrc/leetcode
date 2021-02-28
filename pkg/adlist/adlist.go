package adlist

type List struct {
	len  uint64
	head *listNode
	tail *listNode
}

type listNode struct {
	Val  interface{}
	Prev *listNode
	Next *listNode
}

func NewList() *List {
	return &List{0, nil, nil}
}

func NewNode(val interface{}, prev, next *listNode) *listNode {
	return &listNode{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

func (l *List) AddNodeHead(val interface{}) {
	l.head = NewNode(val, nil, l.head)
	l.len++
}

func (l *List) AddNodeTail(val interface{}) {
	l.tail = NewNode(val, l.tail, nil)
	l.len++
}

func (l *List) RemoveHead() {
}

func (l *List) RemoveTail() {

}

func (l *List) Len() uint64 {
	return l.len
}
