package solution

import (
	"reflect"
	"testing"

	. "github.com/nimrc/leetcode/pkg/linkedlist"
)

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	common := NewList([]int{1, 8, 4, 5})

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				headA: NewList([]int{4}).Append(common),
				headB: NewList([]int{5, 0}).Append(common),
			},
			want: common,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
