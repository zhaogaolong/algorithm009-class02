# 120. 三角形最小路径和

> https://leetcode-cn.com/problems/triangle/

思路分析： 自低而上的 dp 解决思路

推到公式： `dp[i][j] += min(dp[i+1][j], dp[i+1][j+1])`

```go
func minimumTotal(triangle [][]int) int {
    for i := len(triangle)-2; i>= 0; i--{
        for j := range triangle[i]{
            triangle[i][j] += min(
                triangle[i+1][j], triangle[i+1][j+1],
            )
        }
    }
    return triangle[0][0]
}
func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}
```
