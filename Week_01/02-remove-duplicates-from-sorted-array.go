package week01

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++{
		if nums[i] != nums[i-1]{
			count++
			nums[count] = nums[i]
		}
	}
	return count+1
}