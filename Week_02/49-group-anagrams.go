package week02

import "github.com/golang/go/src/sort"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for i := range strs {
		str := strs[i]
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		k := string(b)
		groups[k] = append(groups[k], str)
	}

	res := [][]string{}
	for i := range groups {
		res = append(res, groups[i])
	}

	return res
}
