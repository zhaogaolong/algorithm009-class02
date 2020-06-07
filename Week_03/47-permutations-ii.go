package week03

import "sort"

var res47 [][]int

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	res47 = make([][]int, 0)
	used := make([]bool, len(nums))
	dfs47([]int{}, nums, used)
	return res47

}

func dfs47(temp, nums []int, used []bool) {
	if len(temp) == len(nums) {
		res47 = append(res47, append([]int{}, temp...))
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && used[i-1] {
			continue
		}

		temp = append(temp, nums[i])
		used[i] = true
		dfs47(temp, nums)
		used[i] = false
		temp = temp[:len(temp)-1]
	}
}
