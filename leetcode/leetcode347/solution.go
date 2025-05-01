package leetcode347

func topKFrequent(nums []int, k int) []int {
	freq := make([][]int, len(nums))

	counts := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, ok := counts[num]; !ok {
			counts[num] = 1
			continue
		}

		counts[num] += 1
	}

	for num, f := range counts {
		freq[f-1] = append(freq[f-1], num)
	}

	res := make([]int, 0, k)

	rem := k
	for i := len(freq) - 1; i >= 0; i-- {
		if len(freq[i]) <= rem {
			res = append(res, freq[i]...)
			rem -= len(freq[i])
			continue
		}

		res = append(res, freq[i][:rem]...)
		break
	}

	return res
}
