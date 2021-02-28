package linkedlist

func ExampleNewList() {
	values := []int{2, 4, 3}

	list := NewList(values)

	list.Dump()
	// Output: 2 -> 4 -> 3
}
