package main

import (
	"fmt"

	day1 "github.com/thebadams/advent-of-code-2024/day-1"
	day2 "github.com/thebadams/advent-of-code-2024/day-2"
)

func main() {
	d1p1, d1p2 := day1.Solve("./day-1/input.txt")
	fmt.Printf("Result for day1 part one: %d", d1p1)
	fmt.Printf("Result for day1 part two: %d", d1p2)

	d2p1 := day2.Solve("./day-2/input.txt")
	fmt.Printf("Result for day2 part one: %d", d2p1)
}
