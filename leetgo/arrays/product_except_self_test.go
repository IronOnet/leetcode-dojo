package arrays 

import (
	"testing"
	"reflect"
)

func TestProductExceptSelf(t *testing.T){
	array := []int{1, 2, 3, 4} 
	expected_result := []int{24, 12, 8, 6}  
	actual_result := ProductExceptSelf(array) 

	if !reflect.DeepEqual(expected_result, actual_result){
		t.Errorf("test failure expected %v but got %v", expected_result, actual_result)
	}

	actual_result = ProductExceptSelfII(array) 
	if !reflect.DeepEqual(expected_result, actual_result){
		t.Errorf("test failure expected %v but got %v", expected_result, actual_result)
	}

}