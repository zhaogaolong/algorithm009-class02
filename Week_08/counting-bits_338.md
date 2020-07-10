## 338. 比特位计数

> https://leetcode-cn.com/problems/counting-bits/

```go
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++{
        res[i] = res[i & (i-1)] +1
    }
    return res
}
```
