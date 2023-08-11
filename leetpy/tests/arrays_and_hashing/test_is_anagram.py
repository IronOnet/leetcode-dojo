from arrays_and_hashing.is_anagram import isAnagram

def test_is_anagram(): 
    source, target = "angel", "glean" 
    sourceb, targetb = "couple", "decouple" 
    assert isAnagram(source, target) == True
    assert isAnagram(sourceb, targetb) == False 