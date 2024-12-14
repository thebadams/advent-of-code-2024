package day3

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Running TestSolve Function")
	test1, test2 := 161, 48
	result1, result2 := Solve("./test_input.txt")

	if result1 != test1 {
		t.Errorf("Got unexpected first value. Wanted %d, got %d", test1, result1)
	}

	if result2 != test2 {
		t.Errorf("Got unexpected second value. Wanted %d, got %d", test2, result2)
	}

}
