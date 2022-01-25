package main

import (
	"sort"
	"strings"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

var DIRS = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}

			if dfs(&board, i, j, 0, &word) {
				return true
			}
		}
	}
	return false
}

func dfs(board *[][]byte, i, j, index int, word *string) bool {
	if index == len(*word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[0]) || (*board)[i][j] == '*' || (*word)[index] != (*board)[i][j] {
		return false
	}

	ch := (*board)[i][j]
	(*board)[i][j] = '*'

	for _, dir := range DIRS {
		if dfs(board, i+dir[0], j+dir[1], index+1, word) {
			return true
		}
	}
	(*board)[i][j] = ch
	return false
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := []int{}

	if len(lists) == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		tmp := lists[i]
		for tmp != nil {
			arr = append(arr, tmp.Val)
			tmp = tmp.Next
		}
	}
	if len(arr) == 0 {
		return nil
	}
	sort.Ints(arr)
	start := &ListNode{Val: arr[0]}
	cur := start
	for i := 0; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return start.Next
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}

	mp := make(map[string][]string)
	for _, str := range strs {
		k := getKey(&str)
		mp[k] = append(mp[k], str)
	}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func getKey(s *string) string {
	arr := make([]byte, 26)
	for i := 0; i < len(*s); i++ {
		arr[(*s)[i]-'a']++
	}
	sb := strings.Builder{}
	for _, v := range arr {
		sb.WriteByte(v)
	}
	return sb.String()
}
