package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	type args struct {
		nums     []int
		expected []int
		k        int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "rotate",
			args: args{
				nums:     []int{1, 2, 3, 4, 5, 6, 7},
				expected: []int{5, 6, 7, 1, 2, 3, 4},
				k:        3,
			},
		},

		{
			name: "rotate longer than array length",
			args: args{
				nums:     []int{-1},
				expected: []int{-1},
				k:        2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.args.expected, tt.args.nums)
		})
	}
}
