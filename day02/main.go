package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func isSafe(report []string) bool {
	lastChange := 0
	for index := 1; index < len(report); index++ {
		level, err := strconv.Atoi(report[index])
		if err != nil {
			panic(err)
		}
		prevLevel, err := strconv.Atoi(report[index-1])
		if err != nil {
			panic(err)
		}

		diff := level - prevLevel
		if (diff < -3 || diff > 3) || diff == 0 {
			return false
		}
		if diff > 0 && lastChange < 0 {
			return false
		}
		if diff < 0 && lastChange > 0 {
			return false
		}
		lastChange = diff
	}
	return true
}

func part1(fileContent string) int {
	total := 0
	for _, line := range strings.Split(fileContent, "\n") {
		content := strings.Fields(line)
		if len(content) == 0 {
			continue
		}

		if isSafe(content) {
			total++
		}
	}
	return total
}

func part2(fileContent string) int {
	total := 0
	for _, line := range strings.Split(fileContent, "\n") {
		content := strings.Fields(line)
		if len(content) == 0 {
			continue
		}
		for index := 0; index < len(content); index++ {
			copyContent := make([]string, len(content))
			copy(copyContent, content)

			if index == len(copyContent)-1 {
				copyContent = copyContent[:index]
			} else {
				copyContent = append(copyContent[:index], copyContent[index+1:]...)
			}
			if isSafe(copyContent) {
				total++
				break
			}
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
