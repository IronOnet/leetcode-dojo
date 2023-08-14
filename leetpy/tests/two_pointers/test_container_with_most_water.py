from two_pointers.container_with_most_water import maxArea 

def test_max_area(): 
    input_1 = [1, 8, 6, 2, 5, 4, 8, 3, 7] 
    input_2 = [1, 1] 
    expected_result_1 = 49 
    expected_result_2  = 1 
    
    actual_result_1 = maxArea(input_1) 
    actual_result_2 = maxArea(input_2) 
    assert actual_result_1 == expected_result_1 
    assert actual_result_2 == expected_result_2 