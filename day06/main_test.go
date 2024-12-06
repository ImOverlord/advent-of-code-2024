package main

import (
	"testing"
)

var inputPart1 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	result, _ := part1(inputPart1)

	if result != 41 {
		t.Errorf("part1() = %v, want %v", result, 41)
	}
}

func TestPart2(t *testing.T) {
	_, visits := part1(inputPart1)
	result := part2(inputPart1, visits)

	if result != 6 {
		t.Errorf("part2() = %v, want %v", result, 6)
	}
}
