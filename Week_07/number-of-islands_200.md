## 岛屿数量

> https://leetcode-cn.com/problems/number-of-islands/

```go
func numIslands(grid [][]byte) int {
    counter := 0
    for i := range grid{
        for j := range grid[0]{
            if grid[i][j] == '1'{
                counter++
                dfs(grid, i, j)
            }
        }
    }
    return counter
}

func dfs(grid [][]byte, i, j int){
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]){
        return
    }
    if grid[i][j] != '1'{
        return
    }
    grid[i][j] = '0'
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j-1)
}
```
