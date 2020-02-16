package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/fyibmsd/leetcode/common"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := NewList(3, 4, 2)
	l2 := NewList(4, 6, 5)

	list := addTwoNumbers(l1, l2)

	assert.Equal(t, list.Value(), 807)
}
