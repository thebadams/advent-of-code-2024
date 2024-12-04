package day2

import (
	"fmt"
	"testing"
)

func TestCountSafe(t *testing.T) {
	fmt.Println("Testing CountSafe function")
	input := []bool{true, false, false, false, false, true}
	result := CountSafe(input)
	expected := 2

	if result != expected {
		t.Errorf("Got %d Safe Count, wanted %d", result, expected)
	}
}
