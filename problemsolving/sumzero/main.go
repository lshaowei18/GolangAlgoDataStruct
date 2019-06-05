package sumzero

func sumZero(nums []int) []int {
	ans := []int{}
	//key: 0 - num , value : index of num
	sums := make(map[int]int)

	for i, v := range nums {
		sums[0-v] = i
	}

	for i, v := range nums {
		// to prevent 0 + 0 as the answer
		if v == 0 {
			continue
		}

		//Look for a value that is a valid key
		i2, ok := sums[v]
		if ok {
			// return the values
			ans = append(ans, nums[i], nums[i2])
			return ans
		}
	}

	return ans
}

func multiplePointer(nums []int) []int {
	ans := []int{}

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == 0 {
			ans = append(ans, nums[left], nums[right])
			return ans
		}

		left++
		right--
	}

	return ans
}
