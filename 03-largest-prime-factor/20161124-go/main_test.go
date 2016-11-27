package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLargestPrimeFactorFor(t *testing.T) {
	assert.Equal(t, 29, getLargestPrimeFactorFor(13195))
}
