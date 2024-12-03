package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
	re := regexp.MustCompile(`(?m)mul\(\d*,\d*\)`)
	total := 0

	for _, match := range re.FindAllString(input, -1) {
		line := match[4 : len(match)-1]
		parts := strings.Split(line, ",")
		subTotal := 1

		for _, item := range parts {
			n, _ := strconv.Atoi(item)
			subTotal *= n
		}

		total += subTotal
	}
	return total
}

func part2(input string) int {
	re := regexp.MustCompile(`(?m)don't\(\)|do\(\)`)
	on := true
	prevIndex := 0
	total := 0

	for _, match := range re.FindAllStringIndex(input, -1) {
		action := input[match[0]:match[1]]
		start := match[0]
		end := match[1]
		if action == "don't()" {
			if on == true {
				line := input[prevIndex:start]
				total += part1(line)
				on = false
				prevIndex = end
			}
		} else if action == "do()" {
			if on == false {
				on = true
				prevIndex = end
			}
		}
	}
	if on {
		line := input[prevIndex:]
		total += part1(line)

	}
	return total
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
