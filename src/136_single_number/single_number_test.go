package solution

import "testing"

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		wantN int
	}{
		{
			name:  "example",
			args:  args{[]int{4, 1, 2, 1, 2}},
			wantN: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := singleNumber(tt.args.nums); gotN != tt.wantN {
				t.Errorf("singleNumber() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
