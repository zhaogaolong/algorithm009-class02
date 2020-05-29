package week02

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	res := map[byte]int{}
	for i := range s {
		res[s[i]]++
		res[t[i]]--
	}
	for _, v := range res {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	res := make([]int, 26)
	for i := range s {
		res[int(s[i]-'a')]++
		res[int(t[i]-'a')]++
	}
	for _, v := range res {
		if v != 0 {
			return false
		}
	}

	return true
}
