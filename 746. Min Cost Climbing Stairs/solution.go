func minCostClimbingStairs(cost []int) int {
    var val,res int
    f0 := cost[0]
    f1 := cost[1]
    for i := 2; i <= len(cost); i ++ {
        if i == len(cost){
            val = 0
        } else {
            val = cost[i]
        }
        
        if val + f0 > val + f1{
                res = val + f1
        } else {
                res = val + f0
        }
        fmt.Println("res :",res)
        //顺序不能变，先要覆盖f1
        f0 = f1
        f1 = res
    }
    
    
    return res
}