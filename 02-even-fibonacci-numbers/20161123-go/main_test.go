package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFibonacciUntil(t *testing.T) {
	f := getFibonacciUntil(10)

	assert.Equal(t, []int{1, 2, 3, 5, 8}, f)
}
