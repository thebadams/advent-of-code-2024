package main

import (
	"fmt"

	day1 "github.com/thebadams/advent-of-code-2024/day-1"
)

func main() {
	one, two := day1.Solve("./day-1/input.txt")
	fmt.Printf("Result for day1 part one: %d", one)
	fmt.Printf("Result for day1 part two: %d", two)
}
