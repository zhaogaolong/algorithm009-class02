package week04

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
	i, j, counter := 0, 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			counter++
		}
		j++
	}

	return counter
}
