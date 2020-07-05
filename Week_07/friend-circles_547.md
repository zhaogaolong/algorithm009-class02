## 朋友圈

> https://leetcode-cn.com/problems/friend-circles/

```go
func findCircleNum(m [][]int) int {
    if len(m) == 0{
        return 0
    }
    n := len(m)
    p := NewUnionFind(n)

    for i := range m{
        for j := range m[0]{
            if m[i][j] == 1{
                p.union(i, j)
            }
        }
    }
    return p.count
}

type UnionFind struct{
    count int
    parent []int
}

func NewUnionFind(n int) UnionFind{
    u := UnionFind{ count: n}
    parent := make([]int, n)
    for i := range parent{
        parent[i] = i
    }
    u.parent = parent
    return u
}

func (this *UnionFind)union(p, q int){
    rootP := this.find(p)
    rootQ := this.find(q)
    if rootP == rootQ{
        return
    }
    this.parent[rootP] = rootQ
    this.count--
}

func (this *UnionFind)find(p int) int{
    for p != this.parent[p]{
        this.parent[p] = this.parent[this.parent[p]]
        p = this.parent[p]
    }
    return p
}

```
