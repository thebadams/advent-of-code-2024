package day1

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Testing Solve Function!")
	result := Solve("./test_input.txt")
	expected := 11

	if result != expected {
		t.Errorf("Solved for %d, expected to get %d", result, expected)
	}
}
