package arrays 

func isAnagram(source, target string) bool{
	if len(source) != len(target){
		return false 
	}

	var freq [26]int 

	for idx := 0; idx < len(source); idx++{
		freq[source[idx] - 'a']++ 
		freq[target[idx] - 'a']--
	}

	for idx := 0; idx < len(freq); idx++{
		if freq[idx] != 0{
			return false 
		}
	}
	return true 
}