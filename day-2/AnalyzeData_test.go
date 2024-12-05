package day2

import (
	"fmt"
	"testing"
)

func TestAnalyzeData(t *testing.T) {
	//TODO: Write tests for the different scenarios
	fmt.Println("Testing Analyze Data Function")
	result1 := AnalyzeData([]int{7, 6, 4, 2, 1})
	expected1 := true

	if result1 != expected1 {
		t.Errorf("Unexpected result in first dataset got %v, want %v", result1, expected1)
	}

	result2 := AnalyzeData([]int{1, 2, 7, 8, 9})
	expected2 := false
	if result2 != expected2 {
		t.Errorf("Unexpected result in second dataset, got %v, want %v", result2, expected2)
	}

	result3 := AnalyzeData([]int{9, 7, 6, 2, 1})
	expected3 := false

	if result3 != expected3 {
		t.Errorf("Unexpected result in third dataset, got %v, want %v", result3, expected3)
	}

	result4 := AnalyzeData([]int{1, 3, 2, 4, 5})
	expected4 := false

	if result4 != expected4 {
		t.Errorf("Unexpected result in fourth dataset, got %v, want %v", result4, expected4)
	}

	result5 := AnalyzeData([]int{8, 6, 4, 4, 1})
	expected5 := false
	if result5 != expected5 {
		t.Errorf("Unexpected result in fifth dataset, got %v, want %v", result5, expected5)
	}

	result6 := AnalyzeData([]int{1, 3, 6, 7, 9})
	expected6 := true

	if result6 != expected6 {
		t.Errorf("Unexpected result in sixth dataset, got %v, want %v", result6, expected6)
	}
}
