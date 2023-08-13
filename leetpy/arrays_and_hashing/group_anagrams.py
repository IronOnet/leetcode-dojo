import collections 
from typing import List 

def groupAnagrams(strings : List[str]) -> List[List[str]]: 
    ans = collections.defaultdict(list) 
    
    for s in strings: 
        count = [0] * 26 
        for c in s: 
            count[ord(c) - ord("a")] += 1 
        ans[tuple(count)].append(s) 
    
    return list(ans.values())