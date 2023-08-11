from arrays_and_hashing.two_sum import two_sum

def test_twosums(): 
    target_num = 3 
    array = [1, 2, 4, 6]
    expected_result = [0, 1] 
    actual_result = two_sum(array, target_num) 
    assert expected_result == actual_result