package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type SearchDirection int

const (
	Right SearchDirection = iota
	Left
	Up
	Down
	DiagonalUpLeft
	DiagonalUpRight
	DiagonalDownLeft
	DiagonalDownRight
)

func (s SearchDirection) String() string {
	return [...]string{
		"Right", "Left", "Up", "Down",
		"DiagonalUpLeft", "DiagonalUpRight",
		"DiagonalDownLeft", "DiagonalDownRight",
	}[s]
}

const xmas = "XMAS"

var directionMap = map[SearchDirection][2]int{
	Right:             {1, 0},
	Left:              {-1, 0},
	Up:                {0, -1},
	Down:              {0, 1},
	DiagonalUpLeft:    {-1, -1},
	DiagonalUpRight:   {1, -1},
	DiagonalDownLeft:  {-1, 1},
	DiagonalDownRight: {1, 1},
}

func findChar(board []string, id, x, y int, direction SearchDirection, toFind string) int {
	if safeGet(board, x, y) != toFind[id] {
		return 0
	}
	if id == len(toFind)-1 {
		return 1
	}
	dx, dy := directionMap[direction][0], directionMap[direction][1]
	x, y = x+dx, y+dy
	return findChar(board, id+1, x, y, direction, toFind)
}

func part1(input string) int {
	total := 0
	board := strings.Split(input, "\n")

	for y, line := range board {
		for x := range line {
			for _, direction := range [...]SearchDirection{Left, Right, Up, Down, DiagonalUpLeft, DiagonalUpRight, DiagonalDownLeft, DiagonalDownRight} {
				total += findChar(board, 0, x, y, direction, xmas)
			}
		}
	}
	return total
}

func safeGet(board []string, x, y int) byte {
	if x < 0 || x >= len(board[0]) || y < 0 || y >= len(board) {
		return '.'
	}
	if len(board[y]) == 0 {
		return '.'
	}
	return board[y][x]
}

func findXSAM(board []string, x, y int) int {
	if board[y][x] != 'A' {
		return 0
	}
	line1 := []byte{
		safeGet(board, x-1, y-1),
		safeGet(board, x, y),
		safeGet(board, x+1, y+1),
	}
	line2 := []byte{
		safeGet(board, x-1, y+1),
		safeGet(board, x, y),
		safeGet(board, x+1, y-1),
	}
	if string(line1) != "SAM" && string(line1) != "MAS" {
		return 0
	}
	if string(line2) != "SAM" && string(line2) != "MAS" {
		return 0
	}
	return 1
}

func part2(input string) int {
	total := 0
	board := strings.Split(input, "\n")

	for y, line := range board {
		for x := range line {
			total += findXSAM(board, x, y)
		}
	}

	return total
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
