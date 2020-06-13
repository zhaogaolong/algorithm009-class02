package week04

func canJump(nums []int) bool {
	l := 0
	for i := range nums {
		if i > l {
			return false
		}
		// 这里用到了一个技巧，就是 i + nums[i]
		// 从当前位置 i 向后最大能跳跃的位置
		// 使用 max 得出最大跳跃位置
		l = max(i+nums[i], l)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
