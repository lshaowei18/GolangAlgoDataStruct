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

func (ht *HashTable) Set(key int, val string) {
	ht.keyMap[key] = append(ht.keyMap[key], val)
}

func (ht HashTable) Get(key string) string {
	hash := ht.Hash(key)

	val := ""

	for _, str := range ht.keyMap[hash] {
		if key == str {
			val = str
		}
	}
	return val
}
