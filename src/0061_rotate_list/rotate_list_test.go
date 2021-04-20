package solution

import (
	. "github.com/nimrc/leetcode/pkg/linkedlist"
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head: NewList([]int{0, 1, 2}),
				k:    4,
			},
			want: NewList([]int{2, 0, 1}),
		},
		{
			name: "case 2",
			args: args{
				head: NewList([]int{1}),
				k:    0,
			},
			want: NewList([]int{1}),
		},
		{
			name: "case 3",
			args: args{
				head: NewList([]int{1}),
				k:    1,
			},
			want: NewList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
