# 198. 打家劫舍

> https://leetcode-cn.com/problems/triangle/

解题思路：dp, 自低而上
推到公式：
当前值 = MAX（前两个位置的值+当前值 ， 前一个位置的值）
`dp[i] = max(dp[i-2]+num[i], dp[i-1])`

```go
```
