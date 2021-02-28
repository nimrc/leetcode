package solution

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains duplicate",
			args: args{[]int{1, 2, 3, 1}},
			want: true,
		},
	}

	fns := []struct {
		name string
		fn   func([]int) bool
	}{
		{
			name: "hash map",
			fn:   containsDuplicateMap,
		},
		{
			name: "sort",
			fn:   containsDuplicateSort,
		},
	}

	for _, tt := range tests {
		for _, ff := range fns {
			t.Run(fmt.Sprintf("%s(%s)", tt.name, ff.name), func(t *testing.T) {
				if got := ff.fn(tt.args.nums); got != tt.want {
					t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
