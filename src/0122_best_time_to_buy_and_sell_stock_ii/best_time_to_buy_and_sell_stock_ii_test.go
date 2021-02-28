package solution

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max profit",
			args: args{[]int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "max profit ii",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "max profit iii",
			args: args{[]int{7, 6, 5, 4, 3, 2, 1}},
			want: 0,
		},
	}

	fns := []struct {
		name string
		fn   func([]int) int
	}{
		{
			name: "greedy",
			fn:   maxProfitGreedy,
		},
		{
			name: "dp",
			fn:   maxProfitDP,
		},
	}

	for _, tt := range tests {
		for _, ff := range fns {
			t.Run(fmt.Sprintf("%s(%s)", tt.name, ff.name), func(t *testing.T) {
				if got := ff.fn(tt.args.prices); got != tt.want {
					t.Errorf("maxProfit() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
