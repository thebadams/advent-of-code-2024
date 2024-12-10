package day2

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Testing Solve Day 2")
	result1, result2 := Solve("./test_input.txt")
	expected1 := 2
	expected2 := 4

	if result1 != expected1 {
		t.Errorf("Got %d, Wanted %d for result without dampener", result1, expected1)
	}

	if result2 != expected2 {

		t.Errorf("Got %d, wanted %d for esult with dampener", result2, expected2)

	}
}
