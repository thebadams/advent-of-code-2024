package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func GetNums(input string) ([2]int, error) {
	//expect input to follow pattern of "mul(xxx,xxx)"
	subStr := strings.Split(input, "(")
	if subStr[0] != "mul" {
		return [2]int{}, fmt.Errorf("Malformed DataShape")
	}
	subStr2 := strings.Split(subStr[1], ")")
	numStr := strings.Split(subStr2[0], ",")
	fmt.Println(numStr)
	var num1 int
	var num2 int
	for i, v := range numStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return [2]int{}, err
		}
		if i == 0 {
			num1 = num
		}
		if i == 1 {
			num2 = num
		}
	}

	return [2]int{num1, num2}, nil
}
