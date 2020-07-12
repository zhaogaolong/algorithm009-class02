## 338. 比特位计数

> https://leetcode-cn.com/problems/counting-bits/

思路：
使用位运算中的 `&`

1. `n & (n-1)` 确认是否是 2 的次方数， +1 就是当前有一个位
2. 找到 `n & (n-1)` 的数 +1 就是当前的数字
   a. `10 & 9 = 8`, `bin(8) => 1000` 也就是 1 个 1 在二进制中（`2 ^ 3 =8`），
   b. 通过查找与（&） 数的二进制加 1 就是当前数字的二进制中 1 的个数
   c. 得到推到公式：`n = dp[n & n-1]+1`

```go
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++{
        res[i] = res[i & (i-1)] +1
    }
    return res
}
```
