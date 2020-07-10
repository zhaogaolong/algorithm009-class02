## 190. 颠倒二进制位

> https://leetcode-cn.com/problems/reverse-bits/
> 思路

1. bin --> string --> revert --> int，脱裤子放屁
2. 位运算

```go
func reverseBits(num uint32) uint32 {
    var res uint32
    var power uint32 = 31
    for num != 0{
        res += (num & 1) << power
        num = num >> 1
        power--
    }
    return res
}
```

```go

func reverseBits(num uint32) uint32 {
    var res uint32
    for i :=0; i < 32;i++{
        temp := res << 1 // res * 2
        temp2 := num & 1 // 最后一位是 1 则为 1 ， 类似 n % 2 取模一样

        res = temp | temp2 // 合并
        num >>= 1 //  去除最后一位， 类似：  num * 2
    }
    return res
}

// 终态
func reverseBits(num uint32) uint32 {
    var res uint32
    for i :=0; i < 32;i++{
        res = (res << 1) | (num & 1)
        num >>= 1
    }
    return res
}
```
