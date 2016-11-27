package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumOfSquare(t *testing.T) {
	assert.Equal(t, 385, getSumOfSquareUntil(10))
}

func TestGetSquareOfSumUntil(t *testing.T) {
	assert.Equal(t, 3025, getSquareOfSumUntil(10))
}
