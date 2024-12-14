package main

import (
	"fmt"

	day1 "github.com/thebadams/advent-of-code-2024/day-1"
	day2 "github.com/thebadams/advent-of-code-2024/day-2"
	day3 "github.com/thebadams/advent-of-code-2024/day-3"
)

func main() {
	d1p1, d1p2 := day1.Solve("./day-1/input.txt")
	fmt.Printf("Result for day1 part one: %d\n", d1p1)
	fmt.Printf("Result for day1 part two: %d\n", d1p2)

	d2p1, d2p2 := day2.Solve("./day-2/input.txt")
	fmt.Printf("Result for day2 part one: %d\n", d2p1)
	fmt.Printf("Result for day2 part 2: %d\n", d2p2)

	d3p1, d3p2 := day3.Solve("./day-3/input.txt")
	fmt.Printf("result for day3 part 1: %d\n", d3p1)
	fmt.Printf("result for day3 part 2: %d\n", d3p2)

}
