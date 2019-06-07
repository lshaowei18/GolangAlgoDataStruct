package minsubarraylen

func minSubArrayLen(nums []int, sum int) int {
	minLen := len(nums) + 1
	start := 0
	end := 0
	total := 0

	for start < len(nums) {
		//current window doesn't add to the given sum
		//move the window to right
		if total < sum && end < len(nums) {
			total += nums[end]
			end++
		} else if total >= sum { //Adds up to sum, shrink window
			minLen = min(minLen, end-start)
			total -= nums[start]
			start++
		} else { //total less than required, but reach the end so break
			break
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
