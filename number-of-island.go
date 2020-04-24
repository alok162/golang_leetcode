
// Input:
// 11110
// 11010
// 11000
// 00000

// Output: 1

func numIslands(grid [][]byte) int {
    visited := make([][]bool, len(grid))
    for index := 0; index < len(grid); index ++ {
        visited[index] = make([]bool, len(grid[0]))
    }
    var count int
    
    row := []int{-1, 0, 0, 1}
    col := []int{0, -1, 1, 0}
    for i := 0; i < len(grid); i ++ {
        for j := 0; j < len(grid[0]); j ++ {
            if visited[i][j] == false && string(grid[i][j]) == "1" {
                dfs(i, j, visited, row, col, grid)
                count += 1
            }
        }
    }
    return count
}

func dfs(i int, j int, visited [][]bool, row []int, col []int, grid [][]byte) {
    visited[i][j] = true
    for k:=0;k<4;k++ {
        if isSafe(i + row[k], j + col[k], visited, grid) {
            dfs(i + row[k], j + col[k], visited, row, col, grid)
        }
    }
}

func isSafe(i int, j int, visited [][]bool, grid [][]byte) bool{
    if i >=0 && j >=0 && i < len(grid) && j < len(grid[0]) && visited[i][j] == false && string(grid[i][j]) == "1" {
        return true
    }
    return false
}
