package main

import (
	myString "ImOverlord/advent-of-code-2024/internal/strings"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func part1(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		items := strings.Split(line, ":")
		lineTotal, _ := strconv.Atoi(items[0])
		nums := myString.StringToIntArray(strings.Split(strings.Trim(items[1], " "), " "))
		n := len(nums) // Length of the bitmask
		max := 1 << n  // 2^n

		for i := 0; i < max; i++ {
			calcTotal := 0

			for j := 0; j < n; j++ {
				if (i & (1 << j)) != 0 {
					calcTotal *= nums[j]
				} else {
					calcTotal += nums[j]
				}
			}
			if calcTotal == lineTotal {
				total += calcTotal
				break
			}
		}
	}

	return total
}

func GenerateStrings(n int) []string {
	if n == 0 {
		return []string{""}
	}

	result := []string{}
	subStrings := GenerateStrings(n - 1)
	for _, subString := range subStrings {
		for _, char := range "012" {
			result = append(result, string(char)+subString)
		}
	}
	return result
}

func part2(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		items := strings.Split(line, ":")
		lineTotal, _ := strconv.Atoi(items[0])
		nums := myString.StringToIntArray(strings.Split(strings.Trim(items[1], " "), " "))
		operations := GenerateStrings(len(nums))

		for _, operation := range operations {
			calcTotal := 0
			for i, op := range operation {
				if op == '0' {
					calcTotal *= nums[i]
				} else if op == '1' {
					calcTotal += nums[i]
				} else {
					concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", calcTotal, nums[i]))
					calcTotal = concat
				}
			}
			if calcTotal == lineTotal {
				total += calcTotal
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
