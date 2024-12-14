package day3

import "strings"

func ReadInstructions(ins []string) (instructions1, instructions2 []string) {

	enabled := true
	for _, v := range ins {
		if v == "don't()" {
			enabled = false
		}
		if v == "do()" {
			enabled = true
		}
		if strings.Contains(v, "mul(") && enabled {
			instructions2 = append(instructions2, v)
		}
		if strings.Contains(v, "mul(") {
			instructions1 = append(instructions1, v)
		}
	}
	return instructions1, instructions2

}
