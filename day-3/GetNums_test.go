package day3

import "testing"

func TestGetNums(t *testing.T) {
	result, err := GetNums("mul(2,4)")
	expected := [2]int{2, 4}
	if err != nil {
		t.Errorf("Error Getting Nums: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Wrong length result, got %d, wanted 2", len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Mismatched values in position %d, wanted %d, got %d", i, expected[i], result[i])
		}
	}
}
