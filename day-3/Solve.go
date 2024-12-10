package day3

import "fmt"

func Solve(path string) int {
	data := readFile(path)
	validInstructions := []string{}
	for _, d := range data {
		valid := UncorruptData(d)
		validInstructions = append(validInstructions, valid...)
	}
	nums := [][2]int{}
	for _, v := range validInstructions {
		numbers, err := GetNums(v)
		if err != nil {
			fmt.Printf("Error! %v", err)
		}
		nums = append(nums, numbers)
	}
	sum := 0

	for _, n := range nums {
		mult := Multiply(n)
		sum = sum + mult
	}
	return sum
}
