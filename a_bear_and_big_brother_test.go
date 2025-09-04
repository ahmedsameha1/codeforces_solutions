package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_A_Bear_and_Big_Brother(t *testing.T){
	assert.Equal(t, 2, A_Bear_and_Big_Brother(4, 7))
	assert.Equal(t, 3, A_Bear_and_Big_Brother(4, 9))
	assert.Equal(t, 1, A_Bear_and_Big_Brother(1, 1))
}