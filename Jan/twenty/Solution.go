package twenty

import "sort"

func twoSum(nums []int, target int) []int {
	num_map := make(map[int]int)
	res := []int{}
	for i, num := range nums {
		cur_target := target - num
		if index, ok := num_map[cur_target]; ok {
			res = append(res, i, index)
			return res
		}
		num_map[num] = i
	}
	return res
}

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	m := len(piles)
	sort.Ints(piles)
	l := 1
	r := piles[m-1]

	for l < r {
		mid := l + (r-l)/2
		consumeHour := getHours(piles, mid)
		if consumeHour <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func getHours(piles []int, hour int) int {
	consumeHour := 0
	for _, pile := range piles {
		consumeHour += pile / hour
		if consumeHour%2 != 0 {
			consumeHour++
		}
	}
	return consumeHour
}
