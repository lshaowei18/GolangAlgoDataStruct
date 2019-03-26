package hashtable

type HashTable struct {
	keyMap [53][]string
}

func (ht HashTable) Hash(key string) int {
	PRIME := 31
	total := 0
	for _, c := range key {
		v := int(c) - 96
		total = (total*PRIME + v) % len(ht.keyMap)
	}

	return total
}
