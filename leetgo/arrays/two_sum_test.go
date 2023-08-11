package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	expected_a := []int{0, 1}
	array_a := []int{1, 2, 4, 5}
	target_a := 3
	actual_a := twoSum(array_a, target_a)
	if !reflect.DeepEqual(actual_a, expected_a) {
		t.Errorf("test failure expected %v but got %v", expected_a, actual_a)
	}

	target_b := 4
	array_b := []int{3, 3, 8, 7}
	actual_b := twoSum(array_b, target_b)
	if actual_b != nil {
		t.Errorf("test failure expected %v but got %v", nil, actual_b)
	}
}
