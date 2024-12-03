package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func GetNums(line string) [2]int {
	subStrings := strings.SplitAfter(line, " ")
	num1str := strings.Trim(subStrings[0], " ")
	num2str := strings.Trim(subStrings[len(subStrings)-1], " ")

	num1, err := strconv.Atoi(num1str)
	if err != nil {
		fmt.Printf("Error converting string to integer: %v", err)
	}
	num2, err := strconv.Atoi(num2str)
	if err != nil {
		fmt.Printf("Error converting string to integer: %v", err)
	}

	nums := [2]int{num1, num2}
	return nums
}
