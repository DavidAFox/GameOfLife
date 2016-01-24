package gameoflife

import (
	"strconv"
	"testing"
)

var seed = [][][]int{
	{
		{0, 0, 0},
		{0, 1, 1},
		{0, 1, 0},
	},
	{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	},
	{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	},
}
var expected = [][][]int{
	{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	},
	{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	},
	{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	},
}

func TestNextGen(t *testing.T) {
	for i := 0; i < len(seed); i++ {
		game := NewGame(seed[i])
		game.NextGen()
		result := game.Grid()
		if !sliceOfSlicesEqual(result, expected[i]) {
			t.Error("Grids dont match got: \n", displayGrid(result), "\n expected: \n", displayGrid(expected[i]))
		}
	}
}

func sliceOfSlicesEqual(first, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if len(first[i]) != len(second[i]) {
			return false
		}
		for y := 0; y < len(first[i]); y++ {
			if first[i][y] != second[i][y] {
				return false
			}
		}
	}
	return true
}

func displayGrid(grid [][]int) string {
	var result string
	for i := 0; i < len(grid); i++ {
		for y := 0; y < len(grid[i]); y++ {
			result += strconv.Itoa(grid[i][y])
		}
		result += "\n"
	}
	return result
}
