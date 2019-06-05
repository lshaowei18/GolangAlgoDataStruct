package same

import "math"

func same(nums []int, squares []int) bool {
	//Frequency of values must be the same, so length must be same
	if len(nums) != len(squares) {
		return false
	}

	squareM := make(map[int]int)

	numM := make(map[int]int)

	for _, s := range squares {
		squareM[s]++
	}

	for _, n := range nums {
		numM[n]++
	}

	for k, n := range numM {
		//Calculate the power of key
		p := int(math.Pow(float64(k), 2))

		// values in squareM should be the same as n
		if squareM[p] != n {
			return false
		}
	}

	return true

}
