package arrays 

func groupAnagrams(strings []string) [][]string{
	anagramMap := make(map[[26]int][]string)
	for _, s := range strings{
		var count [26]int 
		for _, c := range s{
			count[c - 'a']++ 
		}
		anagramMap[count] = append(anagramMap[count], s)
	}

	result := make([][]string, len(anagramMap)) 
	idx := 0 
	for _, v := range anagramMap{
		result[idx] = v 
		idx++
	}
	return result 
}

func groupAnagramsII(strings []string) [][]string{
	anagramMap := make(map[[26]int][]string)
	for _, s := range strings{
		var count [26]int  
		for _, c := range s{
			count[c - 'a']++
		}
		anagramMap[count] = append(anagramMap[count], s)
	}
	result := make([][]string, len(anagramMap)) 
	idx := 0 
	for _, v := range anagramMap{
		result[idx] = v 
		idx++ 
	}
	return result 
}