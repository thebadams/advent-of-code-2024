package day1

import (
	"fmt"
	"testing"
)

func TestSumNums(t *testing.T) {
	fmt.Println("Testing SumNums Function!")
	testNums := []int{2, 1, 0, 1, 2, 5}
	result := SumNums(testNums)
	expected := 11

	if result != expected {
		t.Errorf("Sums equalled %d, expected to get %d", result, expected)
	}
}
