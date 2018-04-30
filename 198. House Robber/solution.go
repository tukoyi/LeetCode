func rob(nums []int) int {
    length := len(nums)
    list := make([]int, length)
    var res int
    
    //处理num长度小于3的情况
    if length == 0 {
        return res
    } else if length < 3 {
        for _, v := range nums{
            if v > res{
                res = v
            }
        }
        return res
    }
    
    //处理nums长度大于4的情况
    //状态方程Dn = max(Dn-2 + nums[n], Dn-1) n > 3, Dn表示rob n个房间得到的最大收益
    list[0] = nums[0]
    //list里面保存rob n个房间得到的最大收益，对n=2的边界情况进行判断
    if nums[0] > nums[1]{
        list[1] = nums[0]
    } else {
        list[1] = nums[1]
    } 
    
    for i :=2; i < length; i ++ {
        if list[i - 2] + nums[i] > list[i - 1]{
            list[i] = list[i - 2] + nums[i]    
        } else {
            list[i] = list[i - 1]
        }
    }
    
    return list[length - 1]
    
    
    
}