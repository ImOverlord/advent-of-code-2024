package main

import (
	"ImOverlord/advent-of-code-2024/internal/grid"
	"fmt"
	"strconv"

	_ "embed"

	"github.com/golang-collections/collections/set"
)

//go:embed input.txt
var input string

type Direction int

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

func (d Direction) String() string {
	switch d {
	case DirectionUp:
		return "Up"
	case DirectionDown:
		return "Down"
	case DirectionLeft:
		return "Left"
	case DirectionRight:
		return "Right"
	default:
		return "Unknown"
	}
}

var directionMap = map[Direction][2]int{
	DirectionRight: {1, 0},
	DirectionLeft:  {-1, 0},
	DirectionUp:    {0, -1},
	DirectionDown:  {0, 1},
}

var oppositeDirection = map[Direction]Direction{
	DirectionUp:    DirectionDown,
	DirectionDown:  DirectionUp,
	DirectionLeft:  DirectionRight,
	DirectionRight: DirectionLeft,
}

func PrintMap(topoMap grid.Grid, posX, posY int) {
	for y := range topoMap {
		for x := range topoMap[y] {
			if x == posX && y == posY {
				fmt.Print("#")
			} else {
				fmt.Print(string(topoMap[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func walk(topoMap grid.Grid, x, y int, direction Direction, previousHeight int, peak *set.Set) int {
	currentHeight, err := strconv.Atoi(string(grid.Get(topoMap, x, y, '-')))
	if err != nil || currentHeight != previousHeight+1 {
		return 0
	}
	if currentHeight == 9 {
		peak.Insert(fmt.Sprintf("%d%d", x, y))
		return 1
	}

	outcomes := 0
	for _, newDirection := range []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight} {
		if newDirection == oppositeDirection[direction] {
			continue
		}
		speed := directionMap[newDirection]
		newX := x + speed[0]
		newY := y + speed[1]

		outcomes += walk(topoMap, newX, newY, newDirection, currentHeight, peak)
	}
	return outcomes
}

func part1(input string) int {
	total := 0
	topoMap := grid.CreateGrid(input)

	for y := range topoMap {
		for x := range topoMap[y] {
			if topoMap[y][x] != '0' {
				continue
			}
			peak := set.New()
			for _, newDirection := range []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight} {
				speed := directionMap[newDirection]
				newX := x + speed[0]
				newY := y + speed[1]

				walk(topoMap, newX, newY, newDirection, 0, peak)
			}
			total += peak.Len()
		}
	}

	return total
}

func part2(input string) int {
	total := 0
	topoMap := grid.CreateGrid(input)

	for y := range topoMap {
		for x := range topoMap[y] {
			if topoMap[y][x] != '0' {
				continue
			}
			peak := set.New()
			outcomes := 0
			for _, newDirection := range []Direction{DirectionUp, DirectionDown, DirectionLeft, DirectionRight} {
				speed := directionMap[newDirection]
				newX := x + speed[0]
				newY := y + speed[1]
				outcomes += walk(topoMap, newX, newY, newDirection, 0, peak)
			}
			total += outcomes
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
