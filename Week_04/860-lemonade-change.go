package week04

func lemonadeChange(bills []int) bool {
	res := map[int]int{}
	for _, v := range bills {
		switch v {
		case 5:
			res[5]++
		case 10:
			res[10]++
			res[5]--
		case 20:
			if res[10] > 0 {
				res[10]--
				res[5]--
			} else {
				res[5] -= 3
			}
		}
		if res[5] < 0 {
			return false
		}
	}
	return true
}
