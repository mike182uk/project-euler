package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstTriangleNumberWithNfactors(t *testing.T) {
	assert.Equal(t, 6, getFirstTriangleNumberWithNfactors(4))
}

func TestNumOfDivisors(t *testing.T) {
	assert.Equal(t, 6, numOfDivisors(28))
}
