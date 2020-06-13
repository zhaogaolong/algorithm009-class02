package week04

func numIslands(grid [][]byte) int {
	counter := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				counter++
				dfs200(grid, i, j)
			}
		}
	}
	return counter
}



func dfs200(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfs200(grid, i+1, j)
	dfs200(grid, i-1, j)
	dfs200(grid, i, j+1)
	dfs200(grid, i, j-1)
}
