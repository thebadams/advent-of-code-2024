package day1

import (
	"fmt"
	"testing"
)

func TestGetSimilarityScore(t *testing.T) {
	fmt.Println("TESTING GET SIMILARITY SCORES")
	testGroup1 := []int{3, 4, 2, 1, 3, 3}
	testGroup2 := []int{4, 3, 5, 3, 9, 3}
	result := GetSimilarityScore(testGroup1, testGroup2)
	expected := []int{9, 4, 0, 0, 9, 9}

	if len(result) != len(expected) {
		t.Errorf("Lengths of slices do not match, result was length %d, expected was length %d", len(result), len(expected))
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Values in position %d do not match: wanted %d, got %d", i, expected[i], result[i])
		}

	}
}
