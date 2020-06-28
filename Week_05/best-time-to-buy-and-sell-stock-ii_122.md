## 122. 买卖股票的最佳时机 II

> https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

---

问题重点：

1. 交易次数尽可能的多
2. 买和卖不能在同一天

解题思路：

```go
func maxProfit(prices []int) int {
    dp0, dp1 := 0, -1<<63
    for i := range prices{
        old := dp0
        dp0 = max(dp0, dp1 + prices[i])
        dp1 = max(dp1, old - prices[i])
    }
    return dp0
}

func max(a, b int) int{
    if  a > b{
        return a
    }
    return b
}
```
