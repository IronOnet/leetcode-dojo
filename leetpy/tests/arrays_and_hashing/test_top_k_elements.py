from arrays_and_hashing.top_k_elements import top_k_frequent 

def test_top_k_frequent(): 
    nums = [1, 1, 2, 2, 3]
    expected_result = [1, 2] 
    actual_result = top_k_frequent(nums, 2) 
    assert expected_result == actual_result
    
    nums_b = [1] 
    expected_result = [1] 
    actual_result = top_k_frequent(nums_b, 1) 
    assert expected_result == actual_result 
    