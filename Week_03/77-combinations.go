package week03

var res [][]int

func combine(n int, k int) [][]int {
	if k <= 0 {
		return nil
	}
	res = make([][]int, 0)
	dfs77([]int{}, n, k, 1)
	return res
}

func dfs77(temp []int, n, k, start int) {
	if len(temp) == k {
		res = append(res, append([]int{}, temp...))
		return
	}

	for i := start; i <= n; i++ {
		temp = append(temp, i)
		dfs77(temp, n, k, i+1)
		temp = temp[:len(temp)-1]
	}
}
