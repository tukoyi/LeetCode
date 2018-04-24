//找到递推关系式f(n) = f(n-1) + f(n-2)

func climbStairs(n int) int {

    if n < 3 {
        return n
    } 
    
    var res int
    f0 := 1
    f1 := 2
    for i := 3; i <= n; i ++ {
    	//注意这种形式，保留最近两个结果
        res = f0 + f1
        f0 = f1
        f1 = res
    }
    return res
}