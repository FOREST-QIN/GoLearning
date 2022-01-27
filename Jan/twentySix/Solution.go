package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	dic := make(map[int][]int)
	weight := make(map[*TreeNode]int)
	q := []*TreeNode{}
	minIndex := 0
	q = append(q, root)
	weight[root] = 0

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			w := weight[cur]
			dic[w] = append(dic[w], cur.Val)

			if cur.Left != nil {
				weight[cur.Left] = w - 1
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				weight[cur.Right] = w + 1
				q = append(q, cur.Right)
			}
			if w < minIndex {
				minIndex = w
			}
		}
	}

	for dic[minIndex] != nil {
		res = append(res, dic[minIndex])
		minIndex++
	}
	return res
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	if nums[left] > nums[right] {
		mid := int(left + (right/left)/2)
		lRes := binarySearch(nums, left, mid, target)
		rRes := binarySearch(nums, mid+1, right, target)
		if lRes != -1 {
			return lRes
		} else {
			return rRes
		}
	}

	for left <= right {
		mid := int(left + (right/left)/2)
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	dic := make(map[byte]string)
	getMap(&dic)
	dfs(digits, 0, "", &res, &dic)
	return res
}

func dfs(digits string, index int, prefix string, res *[]string, dic *map[byte]string) {
	if index == len(digits) {
		*res = append(*res, string(prefix))
		return
	}

	for _, ch := range (*dic)[digits[index]] {
		prefix = prefix + string(ch)
		dfs(digits, index+1, prefix, res, dic)
		prefix = prefix[:len(prefix)-1]
	}

}

func getMap(dic *map[byte]string) {
	(*dic)['2'] = "abc"
	(*dic)['3'] = "def"
	(*dic)['4'] = "ghi"
	(*dic)['5'] = "jkl"
	(*dic)['6'] = "mno"
	(*dic)['7'] = "pqrs"
	(*dic)['8'] = "tuv"
	(*dic)['9'] = "wxyz"
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := n - 1
	top := 0
	bottom := m - 1

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		right--
		top++
		bottom--
	}

	if left > right || top > bottom {
		return res
	}
	if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	} else {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	m := len(nums)
	res := []int{}
	deque := []int{}

	for i := 0; i < m; i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		deque = append(deque, i)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
