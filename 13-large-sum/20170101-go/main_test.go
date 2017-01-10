package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumStrNums(t *testing.T) {
	assert.Equal(t, big.NewInt(1368), sumStrNums([]string{"123", "456", "789"}))
}

func TestGetFirstNDigitsOf(t *testing.T) {
	assert.Equal(t, big.NewInt(9876), getFirstNDigitsOf(big.NewInt(987654321), 4))
}

func TestSplitAtStep(t *testing.T) {
	assert.Equal(t, []string{"123", "456", "789"}, splitAtStep("123456789", 3))
}

func TestFormatNumberStr(t *testing.T) {
	ns := `
123
456
789
`

	assert.Equal(t, "123456789", formatNumberStr(ns))
}
