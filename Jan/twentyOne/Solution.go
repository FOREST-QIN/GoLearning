package twentyOne

import "container/list"

func main() {

}

func isValid(s string) bool {
	dict := make(map[byte]byte)
	dict[']'] = '['
	dict[')'] = '('
	dict['}'] = '{'

	stack := list.New()
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' || ch == '{' || ch == '[' {
			stack.PushBack(ch)
		} else {
			if stack.Len() == 0 || stack.Back().Value != dict[ch] {
				return false
			}
			stack.Remove(stack.Back())
		}
	}
	return stack.Len() == 0
}

func minRemoveToMakeValid(s string) string {
	if len(s) == 0 {
		return s
	}
	remove_right := make([]byte, 0)
	left := 0
	unmatched := 0

	for i, ch := range s {
		if ch == '(' {
			left++
			unmatched++
		} else if ch == ')' {
			if unmatched == 0 {
				continue
			}
			unmatched--
		}
		remove_right = append(remove_right, s[i])
	}
	validLeft := left - unmatched
	res := make([]byte, 0)
	for _, ch := range remove_right {
		if ch == '(' {
			if validLeft == 0 {
				continue
			}
			validLeft--
		}
		res = append(res, ch)
	}
	return string(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l := (len(nums1) + len(nums2) + 1) / 2
	r := (len(nums1) + len(nums2) + 2) / 2

	left := binarySearch(nums1, 0, nums2, 0, l)
	right := binarySearch(nums1, 0, nums2, 0, r)
	return float64(left+right) / 2.0
}

func binarySearch(one []int, i int, two []int, j int, k int) int {
	if i >= len(one) {
		return two[j+k-1]
	}
	if j >= len(two) {
		return one[i+k-1]
	}
	if k == 1 {
		if one[i] < two[j] {
			return one[i]
		} else {
			return two[j]
		}
	}
	oneRes := int(1e7 + 1)
	twoRes := int(1e7 + 1)

	if i+k/2-1 < len(one) {
		oneRes = one[i+k/2-1]
	}
	if j+k/2-1 < len(two) {
		twoRes = two[j+k/2-1]
	}

	if oneRes <= twoRes {
		return binarySearch(one, i+k/2, two, j, k-k/2)
	} else {
		return binarySearch(one, i, two, j+k/2, k-k/2)
	}
}
