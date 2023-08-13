package arrays 

import (
	"testing" 
	"reflect"

)

func TestGroupAnagrams(t *testing.T){
	anagrams := []string{"angel", "glean", "aam", "ama", "ding", "gind", "test", "set"}
	expected_grouping := [][]string{{"angel", "glean"}, {"aam", "ama"}, {"ding", "gind"}, {"test"}, {"set"}}
	actual_grouping := groupAnagrams(anagrams) 
	if !reflect.DeepEqual(expected_grouping, actual_grouping){
		t.Errorf("test failure expected %v but got %v", expected_grouping, actual_grouping)
	}
}