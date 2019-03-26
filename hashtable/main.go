package hashtable

func Hash(key string, length int) int {
	total := 0
	for _, c := range key {
		v := int(c) - 96
		total = (total + v) % length
	}

	return total
}
