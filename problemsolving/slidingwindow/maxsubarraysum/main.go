package maxsubarraysum

func maxSubArraySum(arr []int, num int) int {
	sum := 0
	//calculate the initiate sum
	for i := 0; i < num; i++ {
		sum += arr[i]
	}

	max := sum

	for i := num; i < len(arr); i++ {
		//Slide here; minus the value of the first index and add the value of last
		sum = sum - arr[i-num]
		sum = sum + arr[i]

		if sum > max {
			max = sum
		}
	}

	return max
}
