package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumOfPrimesBelow(t *testing.T) {
	assert.Equal(t, 17, getSumOfPrimesBelow(10))
}
