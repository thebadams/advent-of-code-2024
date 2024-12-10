package day3

import (
	"fmt"
	"testing"
)

func TestBuildValidInstruction(t *testing.T) {
	fmt.Println("Testing Valid Instruction Build")

	testStr1 := "mul(2,4)"
	res1 := BuildValidInstruction(testStr1)
	expected1 := []string{"mul(2,4)"}
	if len(res1) != len(expected1) {
		t.Errorf("Unexpected result in first test string length: Wanted %d, got %d", len(expected1), len(res1))
	}
	for i := range res1 {
		if res1[i] != expected1[i] {
			t.Errorf("Unexpected value at index %d in first test string: got %s, wanted %s", i, res1[i], expected1[i])
		}
	}

	testStr2 := "mul"
	res2 := BuildValidInstruction(testStr2)
	expected2 := []string{}

	if len(res2) != len(expected2) {

		t.Errorf("Unexpected result in second test string: Wanted length %d, got %d", len(expected2), len(res2))
	}

	testStr3 := "3,7"
	res3 := BuildValidInstruction(testStr3)
	expected3 := []string{}

	if len(res3) != len(expected3) {

		t.Errorf("Unexpected result in third test string: Wanted length %d, got %d", len(expected3), len(res3))
	}

	testStr4 := "mul(32,64"
	res4 := BuildValidInstruction(testStr4)
	expected4 := []string{}

	if len(res4) != len(expected4) {
		t.Errorf("Unexpected result in fourth test string: Wanted %d, got %d", len(expected4), len(res4))
	}

	testStr5 := "(mul(11,8)mul(8,5))"
	res5 := BuildValidInstruction(testStr5)
	expected5 := []string{"mul(11,8)", "mul(8,5)"}

	if len(res5) != len(expected5) {
		t.Errorf("Unexpected result in fifth test string: wanted length %d, go %d", len(expected5), len(res5))
	}

}
