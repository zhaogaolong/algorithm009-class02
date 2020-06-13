package week04

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		r := prices[i] - prices[i-1]
		if r > 0 {
			res += r
		}
	}

	return res
}

/*
	贪心算法

*/
