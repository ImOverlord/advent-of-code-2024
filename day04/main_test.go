package main

import (
	"testing"
)

var inputPart1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	result := part1(inputPart1)

	if result != 18 {
		t.Errorf("part1() = %v, want %v", result, 18)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputPart1)

	if result != 9 {
		t.Errorf("part2() = %v, want %v", result, 9)
	}
}
