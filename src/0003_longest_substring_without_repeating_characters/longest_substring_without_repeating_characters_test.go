package _003_longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				s: "abba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
