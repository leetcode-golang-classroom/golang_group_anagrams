package sol

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	hash := make(map[([26]int)]([]string))
	result := [][]string{}
	var countCharFreq = func(input string) [26]int {
		var charFreq [26]int
		sLen := len(input)
		for pos := 0; pos < sLen; pos++ {
			charFreq[input[pos]-'a']++
		}
		return charFreq
	}
	for _, val := range strs {
		key := countCharFreq(val)
		hash[key] = append(hash[key], val)
	}
	for _, val := range hash {
		result = append(result, val)
	}
	return result
}
