from two_pointers.is_palindrome import is_valid_palindrome, is_valid_palindrome_2 

def test_is_valid_palindrome(): 
    input_1 = "amanaplanacanalpanama" 
    input_2 = "raceacar" 
    expected_result_1 = True 
    expected_result_2 = False 
    actual_result_1 = is_valid_palindrome(input_1) 
    actual_result_2 = is_valid_palindrome(input_2) 
    assert expected_result_1 == actual_result_1 
    assert expected_result_2 == actual_result_2 
    
    actual_result_1 = is_valid_palindrome_2(input_1) 
    actual_result_2 = is_valid_palindrome_2(input_2) 
    assert actual_result_1 == expected_result_1 
    assert actual_result_2 == expected_result_2 