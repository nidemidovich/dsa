package heap_sort

func heapifyUp(nums []int, n int, i int) {
	largest := i

	left := 2*i + 1
	right := 2*i + 2

	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapifyUp(nums, n, largest)
	}
}

func sortArray(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapifyUp(nums, len(nums), i)
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]

		heapifyUp(nums, i, 0)
	}

	return nums
}
