# 零钱兑换专题

## 322. 零钱兑换

> https://leetcode-cn.com/problems/coin-change/

题目要求：最少的硬币数量凑成总金额
解题思路：

1. 递归
2. BFS
3. DP
   - 子问题
   - dp table 定义
     - 和斐波那契、上楼梯差不多 `f(n) = f(n-1) + f(n-2)`
     - `f(n) = min(f(n-k) (k in [1,2,5])) + 1`
   - dp 方程

```go
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp{
        dp[i] = amount+1
    }
    dp[0] = 0
    for i := 1; i <= amount; i++{
        for j := range coins{
            if coins[j] <= i{
                dp[i] = min(dp[i], dp[i - coins[j]] + 1)
            }
        }
    }
    if dp[amount] > amount{
        return -1
    }
    fmt.Println(dp)
    return dp[amount]
}

func min(a, b int) int{
    if a < b {return a}
    return b
}

```
