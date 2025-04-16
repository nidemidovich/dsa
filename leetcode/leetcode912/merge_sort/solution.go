package merge_sort

// Just a simple merge sort. O(n log n) time, O(n) space.
func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	i := 0
	j := 0

	res := make([]int, 0, len(left)+len(right))

	for {
		if i >= len(left) || j >= len(right) {
			break
		}

		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
