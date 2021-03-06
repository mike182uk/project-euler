package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSmallestMultiple(t *testing.T) {
	assert.Equal(t, 12, getSmallestMultiple(2, 3, 4))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 12, lcm(3, 4))
}

func TestGCD(t *testing.T) {
	assert.Equal(t, 6, gcd(24, 54))
}
