package solution

import "container/heap"

type minHeap struct {
	data []int
}

func (m *minHeap) Len() int {
	return len(m.data)
}

func (m *minHeap) Less(i, j int) bool {
	return m.data[i] < m.data[j]
}

func (m *minHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *minHeap) Push(x interface{}) {
	m.data = append(m.data, x.(int))
}

func (m *minHeap) Pop() interface{} {
	old := m.data
	n := m.Len()
	item := old[n-1]
	m.data = m.data[:n-1]

	return item
}

func findKthLargest(nums []int, k int) int {
	mh := &minHeap{}
	heap.Init(mh)

	for _, num := range nums {
		heap.Push(mh, num)
	}

	for mh.Len() > k {
		heap.Pop(mh)
	}

	return heap.Pop(mh).(int)
}
