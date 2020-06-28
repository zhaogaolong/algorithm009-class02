# 最长公共字符串

思路：使用一个 dp 二维数组矩阵，

| i   | k   |
| --- | --- |
| j   | e   |

地推方程：
....w
..w
这种情况就直接去掉 w，搜索之前的子序列更好，

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)
    dp := make([][]int, m+1)
    for i := range dp{
        dp[i] = make([]int, n+1)
    }
    for i :=1; i<=m; i++{
        for j :=1; j <= n;j++{
            // 如果当前位置相等的话，本次就直接 +1 即可
            if text1[i-1] == text2[j-1]{
                dp[i][j] = 1 + dp[i-1][j-1]
            }else{
                dp[i][j] = max(
                    dp[i-1][j], dp[i][j-1],
                )
            }
        }
    }
    return dp[m][n]
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
```
