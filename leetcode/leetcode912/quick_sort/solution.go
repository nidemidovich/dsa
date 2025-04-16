package quick_sort

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, low, high int) {
	if low >= 0 && high >= 0 && low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot)
		quickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[(low+high)/2]

	i := low
	j := high
	for {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return j
}
