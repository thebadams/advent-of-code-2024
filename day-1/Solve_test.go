package day1

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Testing Solve Function!")
	partOne, partTwo := Solve("./test_input.txt")
	partOneExpected := 11
	partTwoExpected := 31

	if partOne != partOneExpected {
		t.Errorf("Solved for %d, expected to get %d", partOne, partOneExpected)
	}

	if partTwo != partTwoExpected {
		t.Errorf("Solved for %d, expected to get %d", partTwo, partTwoExpected)
	}

}
