//beat 35%
func islandPerimeter(grid [][]int) int {
    var perimeter int
    for i := 0; i < len(grid); i ++{
        for j := 0; j <len(grid[i]); j++{
            if grid[i][j] == 1 {
                //如果格子的值为1，判断上下左右是否为1 或者 超出边界
                //如果是，周长加1
                if i -1 < 0 ||  grid[i - 1][j] == 0{perimeter ++}
                if i + 1 > len(grid) - 1 || grid[i + 1][j] == 0{perimeter ++}
                if j -1 < 0 ||  grid[i][j - 1] == 0{perimeter ++}
                if j + 1 > len(grid[i]) - 1 || grid[i][j + 1] == 0{perimeter ++}
                 
            }
        }
    }
    return perimeter
    
}