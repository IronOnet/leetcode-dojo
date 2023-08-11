from typing import List 

def two_sum(array: List[int], target :int) -> List[int]: 
    num_dict = {} 
    result = []
    for i, n in enumerate(array): 
        diff = target - n 
        if diff in num_dict: 
            return [num_dict[diff], i] 
        num_dict[n] = i 