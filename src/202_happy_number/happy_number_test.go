package solution

import "testing"

func TestIsHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is happy number",
			args: args{19},
			want: true,
		},
		{
			name: "is happy number",
			args: args{7},
			want: true,
		},
		{
			name: "is happy number",
			args: args{29},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
