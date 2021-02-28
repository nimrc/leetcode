package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "reverse string",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)

			assert.Equal(t, string(tt.args.s), "olleh")
		})
	}
}
