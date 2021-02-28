package solution

import (
	"container/heap"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestHeap(t *testing.T) {
	values := []int{31, 41, 59, 26, 53, 58, 97, 93, 23}

	mh := &minHeap{}
	heap.Init(mh)

	for _, val := range values {
		heap.Push(mh, val)
	}

	for mh.Len() > 0 {
		v := heap.Pop(mh)

		spew.Dump(v)
	}
}

func TestFindKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
