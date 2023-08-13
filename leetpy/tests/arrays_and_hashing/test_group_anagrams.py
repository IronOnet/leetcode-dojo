from arrays_and_hashing.group_anagrams import groupAnagrams 

def test_group_anagrams():
    anagrams = ["angel", "glean", "aam", "ama", "ding", "gind", "test", "set"] 
    expected_grouping = [["angel", "glean"], ["aam", "ama"], ["ding", "gind"], ["test"], ["set"]] 
    actual_grouping = groupAnagrams(anagrams) 
    print(actual_grouping)
    assert expected_grouping == actual_grouping