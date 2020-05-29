package week02

func twoSum(nums []int, target int) []int {
	res := map[int]int{}
	for i := range nums {
		v := target - nums[i]
		j, ok := res[v]
		if ok {
			return []int{j, i}
		}
		res[nums[i]] = i
	}
	return nil
}
