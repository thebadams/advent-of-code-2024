package day1

import (
	"fmt"
	"testing"
)

func TestGetNums(t *testing.T) {
	fmt.Println("Testing GetNums Function!")
	testString := "3   4"
	result := GetNums(testString)
	expected := [2]int{3, 4}

	if result != expected {
		t.Errorf("GetNums return %v, wanted %v", result, expected)
	}
}
