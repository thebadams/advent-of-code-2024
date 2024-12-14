package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func RunInstructions(instructions []string) int {

	numbers := [][2]int{}

	for _, v := range instructions {
		str1, found := strings.CutSuffix(v, ")")
		if !found {
			fmt.Println("Suffix ')' not found, malformed data")
		}
		str2, found := strings.CutPrefix(str1, "mul(")
		if !found {
			fmt.Println("Prefix 'mul(' not found, malformed data")
		}

		nums := strings.Split(str2, ",")

		//convert to number
		number1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error converting first number: %v", err)
		}
		number2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error converting second number: %v", err)
		}
		numbers = append(numbers, [2]int{number1, number2})

	}
	sum := 0
	for _, v := range numbers {
		mult := v[0] * v[1]
		sum = sum + mult
	}
	return sum

}
