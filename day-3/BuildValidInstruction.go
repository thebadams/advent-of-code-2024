package day3

import (
	"strings"
)

func BuildValidInstruction(input string) []string {

	strBuilder := strings.Builder{}
	validInstructions := []string{}
	chars := strings.Split(input, "")
	if len(chars) < 8 {
		return validInstructions
	}
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
			switch char {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
			default:
				strBuilder.Reset()
			}
		case 5:
			switch char {
			case ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
			default:
				strBuilder.Reset()
			}

		case 6:
			switch char {
			case ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
			default:
				strBuilder.Reset()
			}
		case 7:

			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, strBuilder.String())
					strBuilder.Reset()
				}
			default:
				strBuilder.Reset()
			}
		case 8:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, strBuilder.String())
					strBuilder.Reset()
				}
			default:
				strBuilder.Reset()
			}
		case 9:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, strBuilder.String())
					strBuilder.Reset()
				}
			default:
				strBuilder.Reset()
			}
		case 10:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, strBuilder.String())
					strBuilder.Reset()
				}
			default:
				strBuilder.Reset()
			}
		case 11:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				strBuilder.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, strBuilder.String())
					strBuilder.Reset()
				}
			default:
				strBuilder.Reset()
			}
		}

	}
	return validInstructions
}
