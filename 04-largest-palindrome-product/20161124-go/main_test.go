package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLargestPalindromeFromProductOf(t *testing.T) {
	assert.Equal(t, 9009, getLargestPalindromeFromProductOf(10, 99))
}

func TestReverseInt(t *testing.T) {
	assert.Equal(t, 321, reverseInt(123))
}
