func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    var sum, res int
    nlimit := len(nums)
    
    if nlimit < 3{
        return sum
    } 
    
    res = nums[0] + nums[1] + nums[2]
    
    for i := 0; i <nlimit -2; i ++{
        lo, lh := i + 1, nlimit -1
        for lo < lh{
            sum := nums[i] + nums[lo] + nums[lh]
            if math.Abs(float64(sum - target)) < math.Abs(float64(res - target)){
                res = sum
                //这里不用做lo,lh的移动
                if sum == target{
                    return sum
                }
            }
            if sum > target {
                lh --
            } else {
                lo ++
            }

        } 
    }
    return res
    
}