package day3

import (
	"strings"
)

func BuildValidInstruction(input string) []string {

	mul := strings.Builder{}
	do := strings.Builder{}
	not := strings.Builder{}
	validInstructions := []string{}
	chars := strings.Split(input, "")
	if len(chars) < 8 {
		return validInstructions
	}
	//check length of chars must be length of "mul(x,x)" as minimum
	for _, char := range chars {
		switch mul.Len() {
		case 0:
			if char == "m" {
				mul.WriteString(char)
			}
		case 1:
			if char == "u" {
				mul.WriteString(char)
			}
			if char != "u" {
				mul.Reset()
			}
		case 2:
			if char == "l" {
				mul.WriteString(char)

			}

			if char != "l" {
				mul.Reset()
			}
		case 3:
			if char == "(" {
				mul.WriteString(char)
			}

			if char != "(" {
				mul.Reset()
			}
		case 4:
			switch char {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
			default:
				mul.Reset()
			}
		case 5:
			switch char {
			case ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
			default:
				mul.Reset()
			}

		case 6:
			switch char {
			case ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
			default:
				mul.Reset()
			}
		case 7:

			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, mul.String())
					mul.Reset()
				}
			default:
				mul.Reset()
			}
		case 8:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, mul.String())
					mul.Reset()
				}
			default:
				mul.Reset()
			}
		case 9:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, mul.String())
					mul.Reset()
				}
			default:
				mul.Reset()
			}
		case 10:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, mul.String())
					mul.Reset()
				}
			default:
				mul.Reset()
			}
		case 11:
			switch char {
			case ")", ",", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				mul.WriteString(char)
				if char == ")" {
					validInstructions = append(validInstructions, mul.String())
					mul.Reset()
				}
			default:
				mul.Reset()
			}
		}
		switch do.Len() {
		case 0:
			if char == "d" {
				do.WriteString(char)
			}
		case 1:
			if char == "o" {
				do.WriteString(char)
			} else {
				do.Reset()
			}

		case 2:
			if char == "(" {
				do.WriteString(char)
			} else {
				do.Reset()
			}
		case 3:
			if char == ")" {
				do.WriteString(char)
				validInstructions = append(validInstructions, do.String())
			} else {
				do.Reset()
			}
		}

		switch not.Len() {
		case 0:
			if char == "d" {
				not.WriteString(char)
			}
		case 1:
			if char == "o" {
				not.WriteString(char)
			} else {
				not.Reset()
			}
		case 2:
			if char == "n" {
				not.WriteString(char)
			} else {
				not.Reset()
			}
		case 3:
			if char == "'" {
				not.WriteString(char)
			} else {
				not.Reset()
			}
		case 4:
			if char == "t" {
				not.WriteString(char)
			} else {
				not.Reset()
			}
		case 5:
			if char == "(" {
				not.WriteString(char)
			} else {
				not.Reset()
			}
		case 6:
			if char == ")" {
				not.WriteString(char)
				validInstructions = append(validInstructions, not.String())
			} else {
				not.Reset()
			}

		}

	}
	return validInstructions
}
