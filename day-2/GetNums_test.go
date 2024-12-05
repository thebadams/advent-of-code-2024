package day2

import (
	"fmt"
	"testing"
)

func TestGetNums(t *testing.T) {
	fmt.Println("Testing Get Nums Function (Day 2)")
	result, err := GetNums("7 6 4 2 1")
	if err != nil {
		t.Errorf("Error converting number string: %v", err)
	}
	expected := []int{7, 6, 4, 2, 1}
	if len(result) != len(expected) {
		t.Errorf("Unexpected Length, got %d, wanted %d", len(result), len(expected))
	}
}
