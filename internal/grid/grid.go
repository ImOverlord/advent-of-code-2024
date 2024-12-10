package grid

import (
	"fmt"
	"strings"
)

type Grid [][]rune

func CreateGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	var grid Grid

	for _, line := range lines {
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	return grid
}

func Get(grid Grid, x, y int, empty rune) rune {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return empty
	}
	if len(grid[y]) == 0 {
		return empty
	}
	return grid[y][x]
}

func Print(grid Grid) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Printf("\n")
	}
}
