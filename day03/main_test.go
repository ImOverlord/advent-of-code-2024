package main

import (
	"testing"
)

var inputPart1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPart1(t *testing.T) {
	result := part1(inputPart1)

	if result != 161 {
		t.Errorf("part1() = %v, want %v", result, 161)
	}
}

var inputPart2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart2(t *testing.T) {
	result := part2(inputPart2)

	if result != 48 {
		t.Errorf("part2() = %v, want %v", result, 48)
	}
}
