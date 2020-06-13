package week04

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !isIn(wordList, endWord) {
		return 0
	}

	step := 0
	used := make([]bool, len(wordList))
	queue := []string{beginWord}
	for len(queue) > 0 {
		step++
		length := len(queue)

		for i := 0; i < length; i++ {
			if queue[i] == endWord {
				return step
			}
			for j, n := range wordList {
				if canNext(queue[i], n) && !used[j] {
					used[j] = true
					queue = append(queue, n)
				}
			}
		}
		queue = queue[length:]
	}
	return 0
}

func isIn(strs []string, s string) bool {
	for _, v := range strs {
		if v == s {
			return true
		}
	}
	return false
}

func canNext(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	counter := 0
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter == 1
}
