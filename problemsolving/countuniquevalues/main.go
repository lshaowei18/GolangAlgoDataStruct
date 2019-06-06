package countuniquevalues

func countUniqueValues(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	unique := make(map[int]struct{})

	//Indices of nums
	left := 0
	right := len(nums) - 1

	for left < right {
		//Left value and right value of the respective indices
		lv := nums[left]
		rv := nums[right]

		//We don't have to check if key exists because we are using empty structs
		//As the value don't matter, we don't have to check if key exists
		unique[lv] = struct{}{}
		unique[rv] = struct{}{}

		left++
		right--
	}

	return len(unique)
}
