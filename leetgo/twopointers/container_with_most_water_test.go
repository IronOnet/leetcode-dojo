package twopointers

import (
	"testing"
)

var second_input = []int{1, 1}
var expected_second_result = 1

func TestMaxArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected_result := 49
	actual_result := maxArea(input)
	if expected_result != actual_result {
		t.Errorf("test failure expected %d but got %d", expected_result, actual_result)
	}
	actual_second_result := maxAreaII(second_input)
	if expected_second_result != actual_second_result {
		t.Errorf("test failure expected %d but got %d", expected_second_result, actual_second_result)
	}
}
