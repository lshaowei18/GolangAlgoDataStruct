package issubsequence

func isSubsequence(sub, str string) bool {
	i := 0

	//loop through each character in str
	for _, v := range str {
		//if substring matches, move to the next index
		if rune(sub[i]) == v {
			i++
		}
		//when index is the same as the len of substring, it is a subsequence
		if i == len(sub) {
			return true
		}

	}

	return false
}
