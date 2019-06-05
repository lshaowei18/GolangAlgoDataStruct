package anagram

func validAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	s1Map := make(map[string]int)
	s2Map := make(map[string]int)

	for _, c := range s1 {
		s1Map[string(c)]++
	}

	for _, c := range s2 {
		s2Map[string(c)]++
	}

	for k, count := range s1Map {
		if s2Map[k] != count {
			return false
		}
	}
	return true
}

func alternative(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	s1Map := make(map[string]int)

	for _, c := range s1 {
		s1Map[string(c)]++
	}

	for _, c := range s2 {
		s := string(c)
		if s1Map[s] <= 0 {
			return false
		}
		s1Map[s]--
	}

	return true
}
