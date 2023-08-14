package twopointers 

import (
	"testing"
)

func TestIsValidPalindrome(t *testing.T){
	input_1 := "amanaplanacanalpanama" 
	input_2 := "raceacar"
	expected_1, expected_2 := true, false 
	actual_1, actual_2 := IsValidPalindrome(input_1), IsValidPalindrome(input_2) 
	if expected_1 != actual_1{
		t.Errorf("test error expecing %v but got %v", expected_1, actual_1)  
	}
	if expected_2 != actual_2{
		t.Errorf("test error expecting %v but got %v", expected_2, actual_2)
	}
}