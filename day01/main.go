package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(fileContent string) int {
	var left []int
	var right []int

	for _, line := range strings.Split(fileContent, "\n") {
		content := strings.Fields(line)
		if len(content) == 0 {
			continue
		}
		leftContent, err := strconv.Atoi(content[0])
		if err != nil {
			panic(err)
		}
		rightContent, err := strconv.Atoi(content[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftContent)
		right = append(right, rightContent)
	}

	slices.Sort(left)
	slices.Sort(right)
	total := 0
	for index := range left {
		if left[index] > right[index] {
			total += left[index] - right[index]
		} else {
			total += right[index] - left[index]
		}
	}
	return total
}

func part2(fileContent string) int {
	var left []int
	var right []int

	for _, line := range strings.Split(fileContent, "\n") {
		content := strings.Fields(line)
		if len(content) == 0 {
			continue
		}
		leftContent, err := strconv.Atoi(content[0])
		if err != nil {
			panic(err)
		}
		rightContent, err := strconv.Atoi(content[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftContent)
		right = append(right, rightContent)
	}

	total := 0

	for _, toFind := range left {
		foundCount := 0
		for _, num := range right {
			if toFind == num {
				foundCount++
			}
		}
		total += toFind * foundCount
	}
	return total
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
