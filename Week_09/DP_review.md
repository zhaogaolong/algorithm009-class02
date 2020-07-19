# DP 回顾和总结

DP 和分治没有本质的区别（关键看有无最优子结构）
共性：找出重复子问题

差异：最优子结构 + 中途可以淘汰次优解

## 常见 DP 问题和状态转义方程

### 爬楼梯

> https://leetcode-cn.com/problems/climbing-stairs/

递推公式：`f(n) = f(n-1) + f(n-2)`

```go

// O(2^n)
func f(n int) int {
	if n <= 1{
		return 1
	}
	return f(n-1) + f(n-2)
}


// O(n)
func f(n int) int{
	if n <= 1{
		return 1
	}

	if n not in memory{
		memory[n] = f(n-1) + f(n-2)
	}
}


// O(n)
func f(n int) int{
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i:=2; i<=n; i++{
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}


// O(n), O(1)
func f(n int) int{
	a, b := 1, 1
	for n > 0{
		b = a+b
		a = b-a
		n--
	}
	return a
}
```

### 不同路径问题

> https://leetcode-cn.com/problems/unique-paths/
> 递推公式：`f(x, y) = f(x-1, y) + f(x, y-1)`

```go
// O(2^(xy)) = O(2^n)
func f(x, y int) int{
	if x == 0 || y == 0{ return 0 }
	if x == 1 || y == 1{ return 1 }
	return f(x-1, y) + f(x, y-1)
}

// O(mn)  O(mn)
func f(x, y int) int{
	if x == 0 || y == 0{ return 0 }
	if x == 1 || y == 1{ return 1 }

	for (x, y) not in mem{
		mem[(x, y)] = f(x-1, y) + f(x, y-1)
	}
}

// O(mn)  O(mn)
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp{
        dp[i] = make([]int, n)
    }
    for i := range dp{
        dp[i][0] = 1
    }
    for i := range dp[0]{
        dp[0][i] = 1
    }

    for i :=1; i<m; i++{
        for j:=1; j<n; j++{
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

```

### 打家劫舍

> https://leetcode-cn.com/problems/house-robber/ > https://leetcode-cn.com/problems/house-robber-ii/

递推公式：
`dp[i] 状态定义 max $ of robber of A[0 --> i]`
`dp[i] = max(dp[i-2]+num[i], dp[i-1])`

更实用大部分难得 DP
`dp[i][0]` 的状态定义：`max $ of robbing A[0 - i]` 且没偷 `num[i]`
`dp[i][1]` 的状态定义：`max $ of robbing A[0 - i]` 且偷了 `num[i]`

`dp[i][0] = max(dp[i-1][0], dp[i-1][1])`
`dp[i][1] = dp[i-1][0] + nums[i]`

### 最小路径和

> https://leetcode-cn.com/problems/minimum-path-sum/

状态定义： `dp[i][j] = minPath(A[1-i][i-j])`
转义方程：`dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + A[i][j]`

### 不同路径-II

递推公式：`f(x, y) = f(x-1, y) + f(x, y-1)`

```go
// O(2^(xy)) = O(2^n)
func f(x, y int) int{
	if x == 0 || y == 0{ return 0 }
	if x == 1 || y == 1{ return 1 }
	return f(x-1, y) + f(x, y-1)
}

// O(mn)  O(mn)
func f(x, y int) int{
	if x == 0 || y == 0{ return 0 }
	if x == 1 || y == 1{ return 1 }

	for (x, y) not in mem{
		mem[(x, y)] = f(x-1, y) + f(x, y-1)
	}
}


// O(mn)  O(mn)
func f(x, y int) int{
	dp := make([][]int, x)
	for i := range dp{
		dp[i] = make([]int, y)
	}

	for i :=1; i<len(dp);i++{
		for j:=1; i<len(dp[i]);i++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[x-1][y-1]
}
```

### 股票买卖问题 1-6 秒杀

DP 定义

`dp[i][k][0 or 1]`

- i 代表 天
- k 代表交易次数
- 0 or 1 代表交易或不交易
  总状态：i _ k _ 2 种状态

```go
for 0 <= i <= I:
   for 0 <= k <= K:
       for s in  [0, 1]:
           dp[i][k][s] = max(buy, sell, rest)
```

转义方程：

没有持有股票:
`dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])`
`max(选择 rest, 选择sell )`
解释：今天没股票的两种可能

- 昨天没有持有，今天选择 rest，所以今天还是没有
- 昨天持有股票，今天选择 sell，所以今天没有持有股票

持有股票:
`dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-][0] - prices[i])`
`max(选择 rest, 选择 buy )`
解释：今天持有股票的两种可能

- 昨天持有股票，今天选择 rest，所以今天还是持有股票
- 昨天没有持有，今天选择 buy，所以今天持有股票
