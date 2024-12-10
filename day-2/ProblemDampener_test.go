package day2

import (
	"fmt"
	"testing"
)

func TestProblemDampener(t *testing.T) {
	fmt.Println("Testing Problem Dampener")
	result1, index1 := ProblemDampener([]int{7, 6, 4, 2, 1})

	if result1 != true {
		t.Errorf("Unexpected result in first data set: Got %v, wanted true", result1)
	}

	if index1 != 0 {
		t.Errorf("Unexpected index result in first data set: got %v, want 0", index1)
	}

	result2, index2 := ProblemDampener([]int{1, 2, 7, 8, 9})

	if result2 != false {
		t.Errorf("Unexpected result in first data set: Got %v, wanted false", result2)
	}

	if index2 != 2 {
		t.Errorf("Unexpected index result in second data set: got %v, want 2", index2)
	}

	result3, index3 := ProblemDampener([]int{9, 7, 6, 2, 1})

	if result3 != false {
		t.Errorf("Unexpected result in third data set: got %v, want false", result3)
	}

	if index3 != 3 {
		t.Errorf("Unexpected result in third data set index, got %d, want 3", index3)
	}

	result4, index4 := ProblemDampener([]int{1, 3, 2, 4, 5})

	if result4 != true {
		t.Errorf("Unexpected result in fourth data set, got %v, wanted true", result4)
	}

	if index4 != 0 {
		t.Errorf("Unexpected index result in fourth data set, got %d, wanted 0", index4)
	}

	result5, index5 := ProblemDampener([]int{8, 6, 4, 4, 1})
	if result5 != true {
		t.Errorf("Unexpected result in fifth data set, got %v, wanted true", result5)
	}

	if index5 != 0 {
		t.Errorf("Unexpected index result in fifth data set, got %d, wanted 0", index5)
	}

	result6, index6 := ProblemDampener([]int{1, 3, 6, 7, 9})
	if result6 != true {
		t.Errorf("Unexpected result in sixth data set, got %v, wanted true", result6)
	}

	if index6 != 0 {
		t.Errorf("Unexpected index result in sixth data set, got %d, wanted 0", index6)
	}

}
