package day3

import (
	"strings"
)

func UncorruptData(input string) []string {
	mapToString := func(r rune) rune {
		if r != 'm' && r != 'u' && r != 'l' && r != '(' && r != ')' && r != ',' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' && r != '0' {
			r = ' '
			return r
		}
		return r
	}

	mapResult := strings.Map(mapToString, input)

	mappedRunes := strings.Split(mapResult, " ")
	validRunes := []string{}
	for _, v := range mappedRunes {
		if !strings.Contains(" ", v) {
			validRunes = append(validRunes, v)
		}
	}
	validInstructions := []string{}
	// take validRunes, and loop through each string, pulling out and building a good pattern of "mul(num,num)"
	for _, v := range validRunes {
		validIns := BuildValidInstruction(v)
		validInstructions = append(validInstructions, validIns...)
	}
	return validInstructions
}
