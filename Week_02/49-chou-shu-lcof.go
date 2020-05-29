package week02

func nthUglyNumber(n int) int {
	res := []int{1}
	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		n := min(min(res[i2]*2, res[i3]*3), res[i5]*5)
		res = append(res, n)

		if n == res[i2]*2 {
			i2++
		}
		if n == res[i3]*3 {
			i3++
		}
		if n == res[i5]*5 {
			i5++
		}

	}
	return res[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
