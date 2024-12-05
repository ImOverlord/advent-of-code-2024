package main

import (
	myString "ImOverlord/advent-of-code-2024/internal/strings"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Manual struct {
	First  int
	Second int
}

func ParseManualPage(line string) Manual {
	parsed := strings.Split(line, "|")
	if len(parsed) != 2 {
		panic("Invalid line format")
	}

	first, err := strconv.Atoi(parsed[0])
	if err != nil {
		panic("Invalid first value")
	}

	second, err := strconv.Atoi(parsed[1])
	if err != nil {
		panic("Invalid second value")
	}

	return Manual{First: first, Second: second}
}

func IsPageValid(firstPage, secondPage int, manuals []Manual) bool {
	for _, manual := range manuals {
		if manual.First == secondPage && manual.Second == firstPage {
			return false
		}
	}
	return true
}

func IsUpdateValid(line string, manual []Manual) int {
	pageStr := strings.Split(line, ",")
	pages := myString.StringToIntArray(pageStr)
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			if !IsPageValid(pages[i], pages[j], manual) {
				return -1
			}
		}
	}
	return pages[len(pages)/2]
}

func part1(input string) int {
	total := 0
	check := false
	manual := []Manual{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			check = true
			continue
		}
		if check {
			midItem := IsUpdateValid(line, manual)
			if midItem != -1 {
				total += midItem
			}
		} else {
			manual = append(manual, ParseManualPage(line))
		}
	}
	return total
}

func FixUpdate(line string, manual []Manual) int {
	pageStr := strings.Split(line, ",")
	pages := myString.StringToIntArray(pageStr)

	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			if !IsPageValid(pages[i], pages[j], manual) {
				pages[i], pages[j] = pages[j], pages[i]
				i = -1 // Reset the outer loop
				break
			}
		}
	}
	midIndex := len(pages) / 2
	return pages[midIndex]
}

func part2(input string) int {
	total := 0
	check := false
	manual := []Manual{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			check = true
			continue
		}
		if check {
			if IsUpdateValid(line, manual) == -1 {
				total += FixUpdate(line, manual)
			}
		} else {
			manual = append(manual, ParseManualPage(line))
		}
	}
	return total
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
