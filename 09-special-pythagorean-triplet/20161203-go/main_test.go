package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPythagoreanTriplet(t *testing.T) {
	a, b, c := getPythagoreanTriplet(1, 2)

	assert.Equal(t, 3, a)
	assert.Equal(t, 4, b)
	assert.Equal(t, 5, c)
}
