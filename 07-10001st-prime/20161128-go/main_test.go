package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrimeAtIndex(t *testing.T) {
	assert.Equal(t, 13, getPrimeAtIndex(6))
}

func TestIsPrime(t *testing.T) {
	assert.Equal(t, true, isPrime(2))
	assert.Equal(t, true, isPrime(3))
	assert.Equal(t, false, isPrime(4))
	assert.Equal(t, true, isPrime(5))
	assert.Equal(t, false, isPrime(6))
	assert.Equal(t, true, isPrime(7))
}
