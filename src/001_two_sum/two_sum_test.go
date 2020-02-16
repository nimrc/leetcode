package solution

import "github.com/stretchr/testify/assert"
import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	assert.Equal(t, twoSum(nums, target), []int{0, 1})
}
