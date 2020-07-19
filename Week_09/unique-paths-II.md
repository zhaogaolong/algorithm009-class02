# 不同路径问题

- https://leetcode-cn.com/problems/unique-paths/
- https://leetcode-cn.com/problems/unique-paths-ii/
- https://leetcode-cn.com/problems/minimum-path-sum/

## 不同路径 -1

> https://leetcode-cn.com/problems/unique-paths/

递推公式：`f(x, y) = f(x-1, y) + f(x, y-1)`

二维数组实现

```go
// O(mn), O(mn)
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

一维数组实现

```go
// O(mn), O(n)
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := range dp{
        dp[i] = 1
    }
    for i := 1; i < m; i++{
        for j := 1; j < n; j++{
            dp[j] += dp[j-1]
        }
    }
    return dp[n-1]
}
```

## 路径问题-II

- https://leetcode-cn.com/problems/unique-paths-ii/

问题概要：如果矩阵中有 障碍就绕过它， 0 代表无障碍， 1 代表有障碍

```go
// O(mn), O(1) 因为不需要构造推到公式需要的数据
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // base 起点就是障碍
    if obstacleGrid[0][0] == 1{
        return 0
    }

    obs := false
    obstacleGrid[0][0] = 1
    for i:=1; i<len(obstacleGrid);i++{
        if obstacleGrid[i][0] == 1 {
            obs = true
        }
        obstacleGrid[i][0] = 1
        if obs{
            obstacleGrid[i][0] = 0
        }
    }
    obs = false
    for i:=1; i<len(obstacleGrid[0]); i++{
        if obstacleGrid[0][i] == 1 {
            obs = true
        }
        obstacleGrid[0][i] = 1
        if obs{
            obstacleGrid[0][i] = 0
        }
    }

    for i :=1; i < len(obstacleGrid); i++{
        for j :=1; j < len(obstacleGrid[i]); j++{
            if obstacleGrid[i][j] == 1{
                obstacleGrid[i][j] = 0
            }else{
                obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
            }
        }
    }

    return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
```

### 路径最小和

> https://leetcode-cn.com/problems/minimum-path-sum/

递推公式：`f(x, y) += min(f(x-1, y), f(x, y-1))`

```go
func minPathSum(grid [][]int) int {
    // 初始化第一行
    for i := 1; i < len(grid[0]); i++{
        grid[0][i] += grid[0][i-1]
    }

    // 初始化第一列
    for i := 1; i < len(grid); i++{
        grid[i][0] += grid[i-1][0]
    }
    for i := 1; i < len(grid); i++{
        for j := 1; j < len(grid[i]);j++{
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}
func min (a, b int) int{
    if a < b{return a}
    return b
}
```
