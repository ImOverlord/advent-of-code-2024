package main

import (
	"testing"
)

var inputPart1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPart1(t *testing.T) {
	result := part1(inputPart1)

	if result != 14 {
		t.Errorf("part1() = %v, want %v", result, 14)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputPart1)

	if result != 34 {
		t.Errorf("part2() = %v, want %v", result, 34)
	}
}
