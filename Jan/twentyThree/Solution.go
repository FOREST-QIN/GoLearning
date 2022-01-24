package main

import (
	"math/rand"
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := longestPalindrome("babad")

	print(res)

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}

func findKthLargest(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		index := rand.Intn(r-l+1) + l
		swap(&nums, index, r)

		i := l
		j := r - 1

		for i <= j {
			if nums[i] >= nums[r] {
				i++
			} else if nums[j] < nums[r] {
				j--
			} else {
				swap(&nums, i, j)
				i++
				j--
			}
		}
		swap(&nums, i, r)
		if i == k-1 {
			return nums[i]
		} else if i < k-1 {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	return nums[k-1]
}

func swap(nums *[]int, i, j int) {
	tmp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = tmp
}

func generateParenthesis(n int) []string {
	res := []string{}
	prefix := ""
	dfs(0, 0, n, &prefix, &res)
	return res
}

func dfs(l, r, n int, prefix *string, res *[]string) {
	if l == n && r == n {
		*res = append(*res, *prefix)
		return
	}
	if l < n {
		*prefix += "("
		dfs(l+1, r, n, prefix, res)
		*prefix = (*prefix)[:len(*prefix)-1]
	}
	if r < l {
		*prefix += ")"
		dfs(l, r+1, n, prefix, res)
		*prefix = (*prefix)[:len(*prefix)-1]
	}
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	sort.Strings(strs)
	sb := strings.Builder{}
	start := strs[0]
	end := strs[n-1]
	i := 0
	for i < len(start) && i < len(end) {
		if start[i] == end[i] {
			sb.WriteByte(start[i])
			i++
		} else {
			break
		}
	}
	return sb.String()
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if num > n {
			continue
		}
		num--
		if nums[num] > 0 {
			nums[num] = -nums[num]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			return i + 1
		}
	}
	return n + 1
}

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	n := len(nums)
	index := n - 2
	for index >= 0 && nums[index] >= nums[index+1] {
		index--
	}
	if index < 0 {
		reverse(&nums, 0, n-1)
		return
	}
	for i := n - 1; i > index; i-- {
		if nums[i] > nums[index] {
			swap(&nums, i, index)
			reverse(&nums, index+1, n-1)
			return
		}
	}

}

func reverse(nums *[]int, i, j int) {
	for i <= j {
		swap(nums, i, j)
		i++
		j--
	}
}

func minMeetingRooms(intervals [][]int) int {
	var n = len(intervals)
	begin := []int{}
	end := []int{}
	for i := 0; i < n; i++ {
		begin = append(begin, intervals[i][0])
		end = append(end, intervals[i][0])
	}
	sort.Ints(begin)
	sort.Ints(end)
	count := 0
	res := 0
	i := 0
	j := 0

	for i < n && j < n {
		if begin[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		res = getMax(res, count)
	}
	return res
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func validPalindrome(s string) bool {
	return check(&s, 0, len(s), false)
}

func check(s *string, i, j int, flag bool) bool {
	for i <= j {
		if (*s)[i] == (*s)[j] {
			i++
			j--
		} else {
			if flag {
				return false
			}
			return check(s, i+1, j, true) || check(s, i, j-1, true)
		}
	}
	return true
}

var max int
var i int
var j int

func longestPalindrome(s string) string {

	i = 0
	j = 0
	max = 0

	for i := 0; i < len(s); i++ {
		getLongest(s, i, i+1)
		getLongest(s, i, i)
	}
	return s[i : j+1]
}

func getLongest(s string, l, r int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	if r-l-1 > max {
		max = r - l - 1
		i = l + 1
		j = r - 1
	}
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	stationMap := make(map[int][]int)
	for i, _ := range routes {
		for _, station := range routes[i] {
			if _, ok := stationMap[station]; !ok {
				stationMap[station] = []int{i}
			} else {
				stationMap[station] = append(stationMap[station], i)
			}

		}
	}

	visitedBus := make(map[int]bool)
	visitedStation := make(map[int]bool)
	q := []int{source}
	steps := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curStation := q[0]
			q = q[1:]
			if curStation == target {
				return steps
			}
			for _, bus := range stationMap[curStation] {
				if _, ok := visitedBus[bus]; ok {
					continue
				}
				visitedBus[bus] = true
				for _, station := range routes[bus] {
					if _, ok := visitedStation[station]; ok {
						continue
					}
					visitedStation[station] = true
					q = append(q, station)
				}
			}
		}
		steps++
	}
	return -1
}
