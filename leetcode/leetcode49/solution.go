package leetcode49

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{{""}}
	}

	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	// uint8 - from `0 <= strs[i].length <= 100`. So, in worst case key[i] = 100
	// uint16 - from `1 <= strs.length <= 10^4`. So, in worst case len(value) = 10^4
	groups := make(map[[26]uint8][]uint16)

	for i, str := range strs {
		var charFreq [26]uint8
		for _, char := range str {
			charFreq[char-'a'] += uint8(1)
		}

		groups[charFreq] = append(groups[charFreq], uint16(i))
	}

	res := make([][]string, 0, len(groups))

	for _, indicies := range groups {
		r := make([]string, 0, len(indicies))
		for _, idx := range indicies {
			r = append(r, strs[idx])
		}
		res = append(res, r)
	}

	return res
}
