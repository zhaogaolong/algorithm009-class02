## 121. 买卖股票的最佳时机

> https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

思路：

- `dp0` 代表手里有钱，没股票
- `dp1` 代表手里有股票,(有可能欠钱)
- 因为只能是一次交易（买+买）,所以就用两个 dp 来表示即可

```go
func maxProfit(prices []int) int {
    // dp0 代表卖，剩的钱
    // dp1 代表买，买剩下的价值
    dp0, dp1 := 0, -1 << 63

    for i := range prices{
        // 清不清仓，
        // dp1（手里有股票） + prices[i]
        dp0 = max(dp0, dp1 + prices[i])

        // 加不加仓
        dp1 = max(dp1, 0 - prices[i])
    }
    return dp0
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
```

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

## 309. 最佳买卖股票时机含冷冻期

> https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

问题重点：

1. 不限制交易次数
2. 股票售出后，第二天不能买入股票（冷冻期一天）

```go
func maxProfit(prices []int) int {
    dp0 := 0
    dp1 := -1 << 32
    pre := 0
    for i := range prices{
        tmp := dp0 //  前一天的 dp0
        dp0 = max(dp0, dp1 + prices[i])
        dp1 = max(dp1, pre - prices[i])
        pre = tmp // 前一天的 dp0，下次 pre - 的时候就是下次循环了，也就是前两天的 dp0 了
    }
    return dp0
}

func max(a, b int) int {
    if a  > b{
        return a
    }
    return b
}
```

## 714. 买卖股票的最佳时机含手续费

> https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

题目要点： 每次交易收取一定的手续费，不限次数交易（多挣交易费嘛 ^\_^）
题目分析：

- `dp0` 代表手里有钱，没股票
  - dp 公式：`dp1 = max(dp1, dp1 + prices[i])`
- `dp1` 代表手里有股票, 从前一天的 dp0 - 今天股价 - 手续费
  - dp 公式：`dp1 = max(dp1, dp0 - prices[i] - fee)`
