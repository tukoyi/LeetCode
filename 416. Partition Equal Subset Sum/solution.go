//首先nums元素之和sum为偶数，才可能有解、
//问题转化为nums中的元素是否能相加得到sum / 2
//对任意一个元素要么加要么不加，又转换为01背包问题
//状态转移方程dp[i] = max(dp[i-num],dp[i]),边界条件dp[0]=true,因为元素之和为0总成立
//



func canPartition(nums []int) bool {

    var  sum int
    for _, v := range nums{
        sum += v
    }
    
    //数组之和为偶数才可能有解
    if  sum % 2 != 0{
        return false
    }
   
    sum = sum / 2
    res := make([]bool,  sum + 1)
    res[0] = true
    
    for _, num := range nums{
        for i := sum; i > 0; i -- {
        //必须从后往前滚动，从而保证每次都是上一次循环的历史数据
            if num <= i {
                res[i] = res[i] || res[i-num] 
            }    
        }
    }
    
    return res[sum]
    
    
}

