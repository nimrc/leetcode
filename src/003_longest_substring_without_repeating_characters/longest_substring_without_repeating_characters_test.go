package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "abcabcbb"

	res := lengthOfLongestSubstring(input)

	assert.Equal(t, res, 3)
}
