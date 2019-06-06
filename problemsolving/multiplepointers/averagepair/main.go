package averagepair

func averagePair(arr []int, avg float64) bool {
	if len(arr) == 0 {
		return false
	}
	a := 0
	b := len(arr) - 1
	for a < b {
		if float64(arr[a])+float64(arr[b])/2 == avg {
			return true
		}
		a++
		b--
	}
	return false
}
