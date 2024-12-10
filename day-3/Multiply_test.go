package day3

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	fmt.Println("Testing Multiply")

	result := Multiply([2]int{2, 4})

	if result != 8 {
		t.Errorf("Error multiplying first dataset: Got %d, want 8", result)
	}
}
