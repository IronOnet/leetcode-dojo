def isAnagram(source : str, target : str) -> bool: 
    if len(source) != len(target): 
        return False 
    
    count_source, count_target = {}, {} 
    for i in range(len(source)): 
        count_source[source[i]] = 1 + count_source.get(source[i], 0) 
        count_target[target[i]] = 1 + count_target.get(target[i], 0) 
    return count_source == count_target 