package main

import (
	"testing"
)

var testInput = `2333133121414131402`

func TestPart1(t *testing.T) {
	result := part1(testInput)

	if result != 1928 {
		t.Errorf("part1() = %v, want %v", result, 1928)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)

	if result != 2858 {
		t.Errorf("part2() = %v, want %v", result, 2858)
	}
}
