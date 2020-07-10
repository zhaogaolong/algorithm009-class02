## 191. 位 1 的个数

> https://leetcode-cn.com/problems/number-of-1-bits/

### 解题-1

十进制转二进制方式，每次对 2 取模，如果是 1， count++

```go
func hammingWeight(num uint32) int {
    count := 0
    for num > 0{
        if num % 2 == 1{
            count++
        }
        num /= 2
    }
    return count
}
```

### 解题-2

解题思路：

1. 使用位运算 与（&），每次 n &= n-1, 去掉二进制末尾一位

```go
func hammingWeight(num uint32) int {
    res := 0
    for num != 0{
        res++
        num &= num-1
    }
    return res
}
```

与运算：

- 当当个都为一的话，该位就是 1
- 使用 & 符做，音意（and）

n &= n-1 是去掉二进制的末尾 1
例如:

96 = 1100000
95 = 1011111

96 & 95 = 1000000 = 64

64 = 1000000
63 = 0111111

64 & 63 = 0000000 = 0
