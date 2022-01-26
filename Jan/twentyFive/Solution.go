package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic := make(map[string]bool)
	for _, str := range wordList {
		dic[str] = true
	}
	if dic[endWord] == false {
		return 0
	}
	q1 := []string{beginWord}
	q2 := []string{endWord}
	step := 0
	for len(q1) > 0 && len(q2) > 0 {
		step++
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		size := len(q1)
		for i := 0; i < size; i++ {
			cur := q1[0]
			q1 = q1[1:]

			for j := 0; j < len(cur); j++ {
				for k := 'a'; k <= 'z'; k++ {
					tmp := cur[:j] + string(k) + cur[j+1:]
					if indexOf(q2, tmp) != -1 {
						return step + 1
					}
					if dic[tmp] == false {
						continue
					}
					delete(dic, tmp)
					q1 = append(q1, tmp)
				}
			}

		}

	}
	return 0
}

func indexOf(slice []string, item string) int {
	for i, _ := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func main() {

}
