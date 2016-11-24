package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLargestPrimeFactorFor(t *testing.T) {
	assert.Equal(t, 29, getLargestPrimeFactorFor(13195))
}
