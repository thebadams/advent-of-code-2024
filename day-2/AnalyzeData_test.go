package day2

import (
	"fmt"
	"testing"
)

func TestAnalyzeData(t *testing.T) {
	//TODO: Write tests for the different scenarios
	fmt.Println("Testing Analyze Data Function")
	result1, index1 := AnalyzeData([]int{7, 6, 4, 2, 1})
	expected1 := true

	if result1 != expected1 {
		t.Errorf("Unexpected result in first dataset got %v, want %v", result1, expected1)
	}
	if index1 != 0 {
		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index1)
	}

	result2, index2 := AnalyzeData([]int{1, 2, 7, 8, 9})
	expected2 := false
	if result2 != expected2 {
		t.Errorf("Unexpected result in second dataset, got %v, want %v", result2, expected2)
	}
	if index2 != 2 {

		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index2)
	}

	result3, index3 := AnalyzeData([]int{9, 7, 6, 2, 1})
	expected3 := false

	if result3 != expected3 {
		t.Errorf("Unexpected result in third dataset, got %v, want %v", result3, expected3)
	}

	if index3 != 3 {
		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index3)
	}

	result4, index4 := AnalyzeData([]int{1, 3, 2, 4, 5})
	expected4 := false

	if result4 != expected4 {
		t.Errorf("Unexpected result in fourth dataset, got %v, want %v", result4, expected4)
	}

	if index4 != 2 {
		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index4)
	}

	result5, index5 := AnalyzeData([]int{8, 6, 4, 4, 1})
	expected5 := false
	if result5 != expected5 {
		t.Errorf("Unexpected result in fifth dataset, got %v, want %v", result5, expected5)
	}

	if index5 != 3 {

		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index5)
	}

	result6, index6 := AnalyzeData([]int{1, 3, 6, 7, 9})
	expected6 := true

	if result6 != expected6 {
		t.Errorf("Unexpected result in sixth dataset, got %v, want %v", result6, expected6)
	}

	if index6 != 0 {
		t.Errorf("Unexpected Index Value returned, wanted 0, got %d", index6)
	}
}
