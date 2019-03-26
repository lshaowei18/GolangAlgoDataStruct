package hashtable

func Hash(key string, length int) int {
	PRIME := 31
	total := 0
	for _, c := range key {
		v := int(c) - 96
		total = (total*PRIME + v) % length
	}

	return total
}
