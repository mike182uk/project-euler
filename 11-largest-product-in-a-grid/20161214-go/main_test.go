package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testGrid = `
1 2 3 4
5 6 7 8
8 7 6 5
4 3 2 1
`

func TestGreatestProductOfAdjacentNumbers(t *testing.T) {
	assert.Equal(t, 56, greatestProductOfAdjacentNumbers(testGrid, 2))
}

func TestGetRowGroups(t *testing.T) {
	rowGroups := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{5, 6},
		[]int{6, 7},
		[]int{7, 8},
		[]int{8, 7},
		[]int{7, 6},
		[]int{6, 5},
		[]int{4, 3},
		[]int{3, 2},
		[]int{2, 1},
	}

	assert.Equal(t, rowGroups, getRowGroups(testGrid, 2))
}

func TestGetColumnGroups(t *testing.T) {
	colGroups := [][]int{
		[]int{1, 5},
		[]int{5, 8},
		[]int{8, 4},
		[]int{2, 6},
		[]int{6, 7},
		[]int{7, 3},
		[]int{3, 7},
		[]int{7, 6},
		[]int{6, 2},
		[]int{4, 8},
		[]int{8, 5},
		[]int{5, 1},
	}

	assert.Equal(t, colGroups, getColumnGroups(testGrid, 2))
}

func TestGetDiagonalGroups(t *testing.T) {
	diagGroups := [][]int{
		// forward
		[]int{1, 6},
		[]int{2, 7},
		[]int{3, 8},
		[]int{5, 7},
		[]int{6, 6},
		[]int{7, 5},
		[]int{8, 3},
		[]int{7, 2},
		[]int{6, 1},
		// reverse
		[]int{4, 7},
		[]int{3, 6},
		[]int{2, 5},
		[]int{8, 6},
		[]int{7, 7},
		[]int{6, 8},
		[]int{5, 2},
		[]int{6, 3},
		[]int{7, 4},
	}

	assert.Equal(t, diagGroups, getDiagonalGroups(testGrid, 2))
}

func TestProductOfSlice(t *testing.T) {
	assert.Equal(t, 120, productOfSlice([]int{2, 3, 4, 5}))
}
