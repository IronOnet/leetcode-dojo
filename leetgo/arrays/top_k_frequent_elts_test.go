package arrays

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElts(t *testing.T) {
	array := []int{1, 1, 1, 2, 2, 3}
	expected_top_3 := []int{1, 2}
	actual_top_3 := topKFrequent(array, 2)
	if !reflect.DeepEqual(expected_top_3, actual_top_3) {
		t.Errorf("test failure expected %v but got %v", expected_top_3, actual_top_3)
	}

	expected_top_2 := []int{1}
	actual_top_2 := topKFrequent(array, 1)
	if !reflect.DeepEqual(expected_top_2, actual_top_2) {
		t.Errorf("test failure expected %v but got %v", expected_top_2, actual_top_2)
	}
}
