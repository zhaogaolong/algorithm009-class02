# 198. 打家劫舍

> https://leetcode-cn.com/problems/triangle/

解题思路：dp
推到公式：
当前值 = MAX（前两个位置的值+当前值 ， 前一个位置的值）
`dp[i] = max(dp[i-2]+num[i], dp[i-1])`

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
