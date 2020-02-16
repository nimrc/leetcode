package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	// 342: 2 -> 4 -> 3
	list := NewList(3, 4, 2)

	assert.Equal(t, list.Val, 2)
	assert.Equal(t, list.Next.Val, 4)
	assert.Equal(t, list.Next.Next.Val, 3)
	assert.Nil(t, list.Next.Next.Next)

}

func TestListToInteger(t *testing.T) {
	list := NewList(3, 4, 2)

	assert.Equal(t, list.Value(), 342)
}
