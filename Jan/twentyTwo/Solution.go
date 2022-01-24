package twentyTwo

func subarraySum(nums []int, k int) int {
	res := 0
	sum := 0
	dict := make(map[int]int)
	dict[0] = 1
	for _, num := range nums {
		sum += num
		if count, ok := dict[sum-k]; ok {
			res += count
		}
		if count, ok := dict[sum]; ok {
			dict[sum] = count + 1
		} else {
			dict[sum] = 1
		}

	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 0
	j := 0
	res := 0
	window := make(map[byte]int)

	for j < len(s) {
		ch := s[j]
		j++
		if count, ok := window[ch]; ok {
			window[ch] = count + 1
		} else {
			window[ch] = 1
		}

		for window[ch] > 1 {
			lCh := s[i]
			i++
			window[lCh]--
		}
		if j-i > res {
			res = j - i
		}
	}
	return res
}
