package day3

import (
	"fmt"
	"testing"
)

func TestRunInstructions(t *testing.T) {
	fmt.Println("Running TestRunInstructions function")
	testIns := []string{"mul(2,4)", "mul(8,5)"}
	expected := 48
	result := RunInstructions(testIns)

	if result != expected {
		t.Errorf("Unexpected result; got %d, wanted %d", result, expected)
	}
}
