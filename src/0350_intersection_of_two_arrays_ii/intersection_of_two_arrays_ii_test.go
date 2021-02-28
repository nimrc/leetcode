package solution

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			want: []int{9, 4},
		},
		{
			name: "case",
			args: args{[]int{3, 1, 2}, []int{1, 1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
