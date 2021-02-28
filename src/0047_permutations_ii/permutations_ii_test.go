package solution

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "permute unique",
			args: args{[]int{1, 1, 2}},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "case",
			args: args{[]int{2, 1, 1, 2}},
			want: [][]int{{2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}, {1, 2, 1, 2}, {1, 2, 2, 1}, {1, 1, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
