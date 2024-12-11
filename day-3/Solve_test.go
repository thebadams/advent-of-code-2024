package day3

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Testing Day 3 Solve")

	result := Solve("./test_input2.txt")
	if result != 161 {
		t.Errorf("Unexprected result: Got %d, wanted 161", result)
	}
}
