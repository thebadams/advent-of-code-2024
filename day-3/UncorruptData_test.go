package day3

import (
	"fmt"
	"testing"
)

func TestUncorruptData(t *testing.T) {
	fmt.Println("Running TestUncorruptData Function")
	testData := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	//expect to return a slice of strings with each instruction
	expected := []string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"}
	uncorruptedData := UncorruptData(testData)

	if len(expected) != len(uncorruptedData) {
		t.Errorf("Length Mismatch between Expected and Result: Got %d, wanted %d", len(expected), len(uncorruptedData))
	}

	for i := range expected {
		if expected[i] != uncorruptedData[i] {
			t.Errorf("Unexpected Value at Index %d; got %s, wanted %s", i, uncorruptedData[i], expected[i])
		}
	}
}
