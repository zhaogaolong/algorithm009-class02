## 70. 爬楼梯 变种

https://leetcode-cn.com/problems/climbing-stairs/

### 本题

---

时间复杂度：O(n)
空间复杂度：O(n), 因为有一个长度为 N 的数组

```go
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i :=2; i<= n;i++{
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

时间复杂度：O(n)
空间复杂度：O(1)

```go
func climbStairs(n int) int {
    a, b := 1, 1
    for n > 0{
        b += a
        a = b-a
        n--
    }
    return a
}
```

### 变形-1: 输出每一步

---

思路：
子问题拆分，

```
    n
  /   \
n-1   n-2
```

进行递归，终止条件：n = 0

时间复杂度：O(2^n)
空间复杂度：O(2^n)

```go
func climbStairs(n int) int {
    dfs(n, []int{})
    return 0
}

func dfs(n int, climbs []int){
    if n == 0{
        fmt.Println(climbs)
        return
    }
    if n >= 1{
        dfs(n-1, append(climbs, 1))
    }
    if n >= 2{
        dfs(n-2, append(climbs, 2))
    }
}
```

output:
参数：n = 5

```bash
[1 1 1 1 1]
[1 1 1 2]
[1 1 2 1]
[1 2 1 1]
[1 2 2]
[2 1 1 1]
[2 1 2]
[2 2 1]
```

## 变形-2 按照固定步长度,并输出每次组合

---

`steps := [1,2,3,4,5]`

```go
func climbStairs(n int, steps []int) int {
    dfs(n, []int{}, steps)
    return n
}

func dfs(n int, climbs []int, steps []int){
    if n == 0{
        fmt.Println(climbs)
        return
    }
    for i := range steps{
        v := steps[i]
        if n >= v{
            dfs(n-v, append(climbs, v), steps)
        }
    }
}

```

`n = 5`

output:

```go
[1 1 1 1 1]
[1 1 1 2]
[1 1 2 1]
[1 1 3]
[1 2 1 1]
[1 2 2]
[1 3 1]
[1 4]
[2 1 1 1]
[2 1 2]
[2 2 1]
[2 3]
[3 1 1]
[3 2]
[4 1]
[5]
```
