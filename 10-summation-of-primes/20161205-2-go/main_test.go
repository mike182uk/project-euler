package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumOfPrimesBelow(t *testing.T) {
	assert.Equal(t, 17, getSumOfPrimesBelow(10))
}

func TestGeneratePrimesUntil(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 7}, generatePrimesUntil(10))
}
