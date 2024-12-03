package main

import (
	"testing"
)

var inputPart1 = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	result := part1(inputPart1)

	if result != 2 {
		t.Errorf("part1() = %v, want %v", result, 2)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputPart1)

	if result != 4 {
		t.Errorf("part2() = %v, want %v", result, 4)
	}
}

var inputEdgeCase = `75 77 80 78 80 83 84 87`

func TestPart2EdgeCase(t *testing.T) {
	result := part2(inputEdgeCase)

	if result != 1 {
		t.Errorf("part2() = %v, want %v", result, 1)
	}
}
