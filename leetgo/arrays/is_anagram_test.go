package arrays 

import "testing" 

func TestIsAanagram(t *testing.T){
	sourceA, targetA := "angel", "glean" 
	sourceB, targetB := "couple", "decouple" 
	expected_a, expected_b := true, false 
	actual_anagram_a := isAnagram(sourceA, targetA) 
	actual_anagram_b := isAnagram(sourceB, targetB) 
	
	if actual_anagram_a != expected_a{
		t.Errorf("test failure expected %v but got %v", expected_a, actual_anagram_a  )
	}
	if actual_anagram_b != expected_b{
		t.Errorf("test failure expected %v but got %v", expected_b, actual_anagram_b)
	}
}