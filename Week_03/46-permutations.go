package week03

var res46 [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res46 = make([][]int, 0)
	dfs46([]int{}, nums)
	return res
}

func dfs46(temp, nums []int) {
	if len(temp) == len(nums) {
		res46 = append(res46, append([]int{}, temp...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if conatine(temp, nums[i]) {
			continue
		}
		temp = append(temp, nums[i])
		dfs46(temp, nums)
		temp = temp[:len(temp)-1]
	}
}

func conatine(nums []int, n int) bool {
	for i := range nums {
		if nums[i] == n {
			return true
		}
	}
	return false
}
