//beat 100 %
//每次在上一行的结果上增加一个元素，然后从最后开始往前求和
func getRow(rowIndex int) []int {
    
    if rowIndex == 0{
        res := []int{1}
        return res
    } else if rowIndex == 1{
        res := []int{1, 1}
        return res
    }
    
    res := make([]int, rowIndex + 1)

    res[0] = 1
    res[1] = 1

    for i := 2; i <= rowIndex; i ++{
    	for j := i; j >= 0; j-- {
    		if j == 0||j == i{
    			res[j] = 1
    			continue
    		}
    		res [j] = res[j] + res[j - 1]
    	}
    }

    
    return res
    
}