# 70. 爬楼梯

> https://leetcode-cn.com/problems/climbing-stairs/

解题思路：斐波那契数列 + DP

```go
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i :=2; i <=n; i++{
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```
