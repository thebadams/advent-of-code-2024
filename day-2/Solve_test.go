package day2

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Testing Solve Day 2")
	result := Solve("./test_input.txt")
	expected := 2

	if result != expected {
		t.Errorf("Got %d, Wanted %d", result, expected)
	}
}
