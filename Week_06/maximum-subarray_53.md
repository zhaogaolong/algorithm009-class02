# 53. 最大子序和

> https://leetcode-cn.com/problems/maximum-subarray/

题解： 找出最大的连续子数组之和

解题思路

1. 暴力：n^2 (n-1 和 n 对比， 两 for 嵌套)
2. DP
   - 分治： max_sum(i) = Max(max_sum(i-1), 0 + a[i]
   - 状态数组定义 f[i]
   - DP 方程： f[i] = Max(f[i-1], 0) + a[i], 为什么是 max(n, 0) 是防止数字是负数

```go
func maxSubArray(nums []int) int {
    dp := nums[:]
    /*
    最大子序和：当前元素自身最大，或包含之前元素最大
    dp 公式：dp[i] = max(num[i], dp[i-1]+num[i])
    */
    for i :=1; i <len(nums); i++{
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        // 优化
        // dp[i] = nums[i] + max(nums[i]+dp[i-1], 0)

    }
    return max(dp...)
}


func max(nums ...int) int{
    r := -1 << 63
    for i := range nums{
        if nums[i] > r{
            r = nums[i]
        }
    }
    return r
}
```

终版

```go
func maxSubArray(nums []int) int {
    res := nums[0]
    for i :=1; i <len(nums); i++{
        nums[i] = max(nums[i-1] + nums[i], nums[i])
        res = max(res, nums[i])
    }
    return res
}
```

## 152. 乘积最大子数组

> https://leetcode-cn.com/problems/maximum-product-subarray/

问题：找出最大连续乘积

要点：

- 保留最大值和最小值（负数有可能），如果当前 n[i] 是负数就双方交换
- 为什么要保留最小值呢，因为负负得正，负数的`绝对值`有可能更大

```go
func maxProduct(nums []int) int {
    res := -1 << 63
    imax, imin := 1, 1
    for i := range nums{
        if nums[i] < 0{
            imax, imin = imin, imax
        }
        imax = max(imax * nums[i], nums[i])
        imin = min(imin * nums[i], nums[i])
        res = max(res, imax)
    
    return res
}
```
