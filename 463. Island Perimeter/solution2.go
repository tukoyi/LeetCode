//beat 35%
func islandPerimeter(grid [][]int) int {
    var islands,neighbours int
    for i := 0; i < len(grid); i ++{
        for j := 0; j <len(grid[i]); j++{
            if grid[i][j] == 1 {
                //如果格子的值为1，判断右边和下边是否有相邻的格子为1
                //如果是neighbours ++
                //周长 = 4 * islands - 2 * neighbours
                if  j + 1 < len(grid[i]) && grid[i][j + 1] == 1 {
                     neighbours ++
                }
                if  i + 1 < len(grid) && grid[i + 1][j] == 1 {
                     neighbours ++
                }
            islands ++
                }
        }
    }
    return islands * 4 - neighbours * 2
    
}