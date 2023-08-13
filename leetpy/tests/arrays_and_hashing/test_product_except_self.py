from arrays_and_hashing.product_except_self import productExceptSelf, product_except_self

def test_product_except_self(): 
    array = [1, 2, 3, 4] 
    expected_result = [24, 12, 8, 6] 
    actual_result = productExceptSelf(array) 
    assert expected_result == actual_result
    actual_result = product_except_self(array) 
    assert expected_result == actual_result