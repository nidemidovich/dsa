package leetcode169

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	majority := nums[0]
	counter := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			counter++
		} else {
			counter--
		}

		if counter == 0 {
			majority = nums[i]
			counter++
		}
	}

	return majority
}
