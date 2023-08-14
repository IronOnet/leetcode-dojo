package twopointers 

func maxArea(height []int) int{
	left := 0 
	right := len(height)-1
	result := 0 

	for left < right{
		area := min(height[left], height[right]) * (right - left)
		if area > result{
			result = area 
		}

		if height[left] > height[right]{
			right-- 
		} else{
			left++
		}
	}
	return result 
}

func maxAreaII(heights []int) int{
	left, right := 0, len(heights)-1 
	result := 0 
	for left < right{
		result = max(result, min(heights[left], heights[right]) * (right - left))
		
		if heights[left] > heights[right]{
			right-- 
		} else{
			left++ 
		}
	}
	return result 
}

func min(a, b int) int{
	if a < b{
		return a 
	}
	return b 
}

func max(a, b int) int{
	if a > b{
		return a 
	}
	return b 
}