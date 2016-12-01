package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrimeAtIndex(t *testing.T) {
	assert.Equal(t, 13, getPrimeAtIndex(6))
}
