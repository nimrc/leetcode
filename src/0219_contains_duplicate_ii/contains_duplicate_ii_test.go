package solution

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains nearby duplicate",
			args: args{[]int{1, 2, 3, 1}, 3},
			want: true,
		},
		{
			name: "contains nearby duplicate",
			args: args{[]int{1, 2, 3, 1, 2, 3}, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
