package day3

import (
	"strings"
)

func UncorruptData(input string) ([]string, []string) {
	mapToString1 := func(r rune) rune {
		if r != 'm' && r != 'u' && r != 'l' && r != '(' && r != ')' && r != ',' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' && r != '0' {
			r = ' '
			return r
		}
		return r
	}

	mapToString2 := func(r rune) rune {
		if r != 'm' && r != 'u' && r != 'l' && r != '(' && r != ')' && r != ',' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' && r != '0' && r != 'd' && r != 'o' && r != 'n' && r != '\'' && r != 't' {
			r = ' '
			return r
		}
		return r
	}

	mapResult1 := strings.Map(mapToString1, input)
	mapResult2 := strings.Map(mapToString2, input)

	mappedRunes1 := strings.Split(mapResult1, " ")
	mappedRunes2 := strings.Split(mapResult2, " ")
	validRunes1 := []string{}
	validRunes2 := []string{}
	for _, v := range mappedRunes1 {
		if !strings.Contains(" ", v) {
			validRunes1 = append(validRunes1, v)
		}
	}
	for _, v := range mappedRunes2 {
		if !strings.Contains(" ", v) {
			validRunes2 = append(validRunes2, v)
		}
	}
	validInstructions1 := []string{}
	validInstructions2 := []string{}
	// take validRunes, and loop through each string, pulling out and building a good pattern of "mul(num,num)"
	for _, v := range validRunes1 {
		validIns := BuildValidInstruction(v)
		validInstructions1 = append(validInstructions1, validIns...)
	}
	return validInstructions1, validInstructions2
}
