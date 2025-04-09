package leetcode27

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	if len(nums) == 1 && nums[0] != val {
		return 1
	}

	p1 := 0
	p2 := len(nums) - 1

	for {
		if p1 > p2 {
			break
		}

		if nums[p1] == val {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2 -= 1
		} else {
			p1 += 1
		}
	}

	return p1
}
