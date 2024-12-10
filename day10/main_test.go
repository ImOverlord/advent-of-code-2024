package main

import (
	"testing"
)

var testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	res := 36

	if result != res {
		t.Errorf("part1() = %v, want %v", result, res)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	res := 81

	if result != res {
		t.Errorf("part2() = %v, want %v", result, res)
	}
}
