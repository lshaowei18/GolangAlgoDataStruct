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

//This only works if the slice is sorted
func noMaps(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0

	for i := 1; i < len(nums); i++ {
		//if the values are the same, move on to the next
		if nums[index] == nums[i] {
			continue
		}
		//increase index and assign the value of nums[i] to it
		index++
		nums[index] = nums[i]
	}

	return index + 1
}
