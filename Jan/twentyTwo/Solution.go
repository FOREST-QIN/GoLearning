package twentyTwo

func subarraySum(nums []int, k int) int {

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
