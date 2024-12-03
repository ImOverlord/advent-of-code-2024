package main

import (
	"testing"
)

var inputPart1 = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPart1(t *testing.T) {
	result := part1(inputPart1)

	if result != 11 {
		t.Errorf("part1() = %v, want %v", result, 11)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputPart1)

	if result != 31 {
		t.Errorf("part2() = %v, want %v", result, 31)
	}
}
