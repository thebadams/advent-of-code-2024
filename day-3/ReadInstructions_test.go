package day3

import (
	"fmt"
	"testing"
)

func TestReadInstructions(t *testing.T) {
	fmt.Println("Running TestReadInstructions function")
	testIns := []string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"}
	readIns1, readIns2 := ReadInstructions(testIns)
	expected2 := []string{"mul(2,4)", "mul(8,5)"}
	expected1 := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	if len(readIns2) != len(expected2) {
		t.Errorf("Length Mismatch in readIns vs expected; got %d, wanted %d", len(readIns2), len(expected2))
	}

	for i := range readIns2 {
		if readIns2[i] != expected2[i] {
			t.Errorf("Mismatched Value at index %d, got %s and wanted %s", i, readIns2[i], expected2[i])
		}
	}

	if len(readIns1) != len(expected1) {
		t.Errorf("Length mismatch in ReadIns1, got %d, wanted %d", len(readIns1), len(expected1))
	}

	for i := range readIns1 {

		if readIns1[i] != expected1[i] {
			t.Errorf("Mismatched value at index %d, got %s and wanted %s", i, readIns1[i], expected1[i])
		}
	}
}
