package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToIntArray(input string) []int {
	fields := strings.Fields(input)
	intArray := make([]int, len(fields))
	for i, str := range fields {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting %q to integer: %v\n", str, err)
			panic("ConvertToIntArray")
		}
		intArray[i] = num
	}

	return intArray
}
