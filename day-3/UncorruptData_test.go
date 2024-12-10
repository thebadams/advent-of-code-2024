package day3

import (
	"fmt"
	"testing"
)

func TestUncorruptData(t *testing.T) {
	fmt.Println("Testing Uncorrupt Data Function")
	data := UncorruptData("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	expectedData := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	if len(data) != len(expectedData) {
		t.Errorf("Unexpected result in data, wanted length %d, got length %d", len(expectedData), len(data))
	}
}
