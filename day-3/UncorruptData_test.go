package day3

import "testing"

func TestUncorruptData(t *testing.T) {
	UncorruptData("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	t.Error("Test Ran, but function not finished")
}
