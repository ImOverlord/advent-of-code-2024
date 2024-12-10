package main

import (
	"ImOverlord/advent-of-code-2024/internal/grid"
	_ "embed"
	"fmt"
	"math"

	"github.com/golang-collections/collections/set"
)

//go:embed input.txt
var input string

type Antena struct {
	Antena rune
	X      int
	Y      int
}

func distance(p1, p2 Antena) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2))
}

func isColinear(a, b Antena, x, y int) bool {
	area := a.X*(b.Y-y) + b.X*(y-a.Y) + x*(a.Y-b.Y)
	return area == 0
}

func part1(input string) int {
	board := grid.CreateGrid(input)
	antenas := []Antena{}

	for y := range board {
		for x := range board[y] {

			item := grid.Get(board, x, y, '.')
			if item != '.' {
				antenas = append(antenas, Antena{X: x, Y: y, Antena: item})
			}
		}
	}
	nodes := set.New()
	for _, antena := range antenas {
		for _, checkAntena := range antenas {
			if antena == checkAntena || antena.Antena != checkAntena.Antena {
				continue
			}
			vectorX := checkAntena.X - antena.X
			vectorY := checkAntena.Y - antena.Y

			thirdPoint := Antena{
				X: checkAntena.X + vectorX,
				Y: checkAntena.Y + vectorY,
			}
			if grid.Get(board, thirdPoint.X, thirdPoint.Y, '@') != '@' {
				nodes.Insert(thirdPoint)
			}
		}
	}
	return nodes.Len()
}

func part2(input string) int {
	board := grid.CreateGrid(input)
	antenas := []Antena{}

	for y := range board {
		for x := range board[y] {

			item := grid.Get(board, x, y, '.')
			if item != '.' {
				antenas = append(antenas, Antena{X: x, Y: y, Antena: item})
			}
		}
	}
	nodes := set.New()
	for y := range board {
		for x := range board[y] {
			for _, antena := range antenas {
				for _, checkAntena := range antenas {
					if antena == checkAntena || antena.Antena != checkAntena.Antena {
						continue
					}
					if isColinear(antena, checkAntena, x, y) {
						nodes.Insert(Antena{X: x, Y: y, Antena: '#'})
					}
				}
			}
		}
	}
	return nodes.Len()
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
