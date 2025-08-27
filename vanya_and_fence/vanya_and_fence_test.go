package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVanyaAndFence(t *testing.T) {
	result := VanyaAndFence(7, []int{4, 5, 14})
	assert.Equal(t, 4, result)
	result = VanyaAndFence(1, []int{1, 1, 1, 1, 1, 1})
	assert.Equal(t, 6, result)
	result = VanyaAndFence(5, []int{7, 6, 8, 9, 10, 5})
	assert.Equal(t, 11, result)
}