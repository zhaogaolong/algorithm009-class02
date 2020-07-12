# 最长公共字符串

思路：使用一个 dp 二维数组矩阵，

| -   | a   | b   | c   | d   | e   |
| --- | --- | --- | --- | --- | --- |
| a   | 0   | 1   | 2   | 3   | 4   |
| c   | 0   | 1   | 2   | 3   | 4   |
| e   | 0   | 1   | 2   | 3   | 4   |

地推方程：

子问题依赖：

1. 如果 `dp[i][j]` 字母相同， 就是它的斜上角的值 +1，为啥捏：
   - 根据分治思想，如果字符串末尾相同： `abcd` 和 `acd`, 那么直接删除相同字母 `d` 对比 `abc` 和 `ac` 形成新的子问题，对应矩阵操作：👇
   - `dp[i][j] = 1 + dp[i-1][j-1]`
2. 如果 `dp[i][j]` 字母不同，正常继承即可，取矩阵上面和左面中的最大值
   - `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`

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
