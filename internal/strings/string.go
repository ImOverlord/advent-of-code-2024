package strings

import (
	"strconv"
)

func StringToIntArray(stringArray []string) []int {
	intArray := make([]int, len(stringArray))
	for i, str := range stringArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArray[i] = num
	}
	return intArray
}
