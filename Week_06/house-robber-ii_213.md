# 213. 打家劫舍 II

> https://leetcode-cn.com/problems/house-robber-ii/

思路：分为两步

- 从第一个开始, 直到最后前两个（`n-2`）
- 从第二个开始，直到最后一个（`n-1`）
- `return max(dp1[n-1], dp2[n-2])`

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 0{
        return 0
    }else if n == 1{
        return nums[0]
    }

    dp1 := make([]int, n)
    dp1[0] = 0 // 第一个不抢
    dp1[1] = nums[1]
    for i :=2; i <n; i++{
        dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
    }

    dp2 := make([]int, n)
    dp2[0], dp2[1] = nums[0], max(nums[0], nums[1])
    // i < n-1 是最后一个不抢，和上面的那个第一个不抢刚好隔开
    for i :=2; i < n-1;i++{
        dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
    }

    return max(dp1[n-1], dp2[n-2])
}
func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
```
