package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Constants for directions
const (
	DirectionUp    = '^'
	DirectionRight = '>'
	DirectionDown  = 'v'
	DirectionLeft  = '<'
)

var directionMap = map[rune][2]int{
	DirectionRight: {1, 0},
	DirectionLeft:  {-1, 0},
	DirectionUp:    {0, -1},
	DirectionDown:  {0, 1},
}

func nextDirection(current rune) rune {
	switch current {
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	default:
		return DirectionUp
	}
}

func createBoard(input string) [][]rune {
	lines := strings.Split(input, "\n")
	var grid [][]rune

	for _, line := range lines {
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	return grid
}

func safeGet(board [][]rune, x, y int) rune {
	if x < 0 || x >= len(board[0]) || y < 0 || y >= len(board) {
		return '@'
	}
	if len(board[y]) == 0 {
		return '@'
	}
	return board[y][x]
}

func findStartingPos(board [][]rune) (x int, y int, dir rune) {
	for y, row := range board {
		for x, char := range row {
			switch char {
			case DirectionUp, DirectionRight, DirectionDown, DirectionLeft:
				return x, y, char
			}
		}
	}
	panic("Start position not found")
}

func part1(input string) (int, []Visit) {
	total := 0
	board := createBoard(input)
	x, y, direction := findStartingPos(board)
	visit := []Visit{}
	for {
		dx, dy := directionMap[direction][0], directionMap[direction][1]
		nextBlock := safeGet(board, x+dx, y+dy)
		if nextBlock == '#' {
			direction = nextDirection(direction)
		} else if nextBlock == '@' {
			break
		} else {
			x, y = x+dx, y+dy
			if board[y][x] != 'X' {
				visit = append(visit, Visit{X: x, Y: y, Dir: direction})
				total += 1
			}
			board[y][x] = 'X'
		}
	}

	return total, visit
}

type Visit struct {
	X   int
	Y   int
	Dir rune
}

func hasVisitedAlready(x, y int, dir rune, visits []Visit) bool {
	for _, visit := range visits {
		if visit.X == x && visit.Y == y && visit.Dir == dir {
			return true
		}
	}
	return false
}

func part2(input string, prevVisits []Visit) int {
	total := 0
	board := createBoard(input)

	for _, prevVisit := range prevVisits {
		dx, dy := prevVisit.X, prevVisit.Y
		if board[dy][dx] != '.' {
			continue
		}
		newBoard := createBoard(input)
		visits := []Visit{}
		newBoard[dy][dx] = '0'
		x, y, direction := findStartingPos(newBoard)

		for {
			dx, dy := directionMap[direction][0], directionMap[direction][1]
			nextBlock := safeGet(newBoard, x+dx, y+dy)
			if nextBlock == '#' || nextBlock == '0' {
				direction = nextDirection(direction)
			} else if nextBlock == '@' {
				break
			} else if hasVisitedAlready(x, y, direction, visits) {
				total = total + 1
				break
			} else {
				visits = append(visits, Visit{X: x, Y: y, Dir: direction})
				x, y = x+dx, y+dy
				newBoard[y][x] = 'X'
			}
		}
	}

	return total
}

func main() {
	part1Result, visit := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input, visit)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
