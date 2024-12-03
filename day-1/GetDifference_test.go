package day1

import (
	"fmt"
	"testing"
)

func TestGetDifference(t *testing.T) {
	fmt.Println("Testing GetDifference Function!!")
	input1 := []int{1, 2, 3, 3, 3, 4}
	input2 := []int{3, 3, 3, 4, 5, 9}
	result := GetDifference(input1, input2)
	expected := []int{2, 1, 0, 1, 2, 5}

	if len(result) != len(expected) {
		t.Errorf("Lengths do not match: expected was length %d, result yielded length %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Wrong value in position %d: got %d, wanted %d", i, result[i], expected[i])
		}
	}
}
