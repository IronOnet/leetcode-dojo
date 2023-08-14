package twopointers 

func IsValidPalindrome(str string) bool{
	left, right := 0, len(str)-1 

	for i := 0; i < len(str); i++{
		if str[left] != str[right]{
			return false 
		}
		left++ 
		right-- 
	}
	return true 
}