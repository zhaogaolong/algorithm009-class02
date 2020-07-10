## 231. 2 的幂

> https://leetcode-cn.com/problems/power-of-two/

```go
func isPowerOfTwo(n int) bool {
    if n != 0{
        return false
    }
    return n & (n-1) == 0
}
j
func isPowerOfTwo(n int) bool {
    return n != 0 && n & (n-1) == 0
}
```
