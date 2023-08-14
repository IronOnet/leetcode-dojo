def maxArea(heights : int) -> int: 
    left, right = 0, len(heights) -1 
    result = 0 
    while left < right: 
        result = max(result, min(heights[left], heights[right])* (right-left))
         
        if heights[left] < heights[right]: 
            left += 1 
        elif heights[right] <= heights[left]: 
            right -= 1 
    return result 