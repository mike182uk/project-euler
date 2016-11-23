package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFibonacciUntil(t *testing.T) {
	f := GetFibonacciUntil(10)

	assert.Equal(t, []int{1, 2, 3, 5, 8}, f)
}
