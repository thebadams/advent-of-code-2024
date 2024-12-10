package day3

import (
	"fmt"
	"strings"
)

func UncorruptData(input string) {
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
	// take validRunes, and loop through each string, pulling out and building a good pattern of "mul(num,num)"
	for _, v := range validRunes {
		strBuilder := strings.Builder{}
		chars := strings.Split(v, "")
		//check length of chars must be length of "mul(x,x)" as minimum
		for _, char := range chars {
			switch strBuilder.Len() {
			case 0:
				if char == "m" {
					strBuilder.WriteString(char)
				}
			case 1:
				if char == "u" {
					strBuilder.WriteString(char)
				}
				if char != "u" {
					strBuilder.Reset()
				}
			case 2:
				if char == "l" {
					strBuilder.WriteString(char)

				}

				if char != "l" {
					strBuilder.Reset()
				}
			case 3:
				if char == "(" {
					strBuilder.WriteString(char)
				}

				if char != "(" {
					strBuilder.Reset()
				}
			case 4:
				if char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9" {
					strBuilder.WriteString(char)
				}
				if char != "1" && char != "2" && char != "3" && char != "4" && char != "5" && char != "6" && char != "7" && char != "8" && char != "9" {
					strBuilder.Reset()
				}

			}
		}
		// check last character of string for validity
		fmt.Println(strBuilder.String())
	}
}
