//用two point方法，转换为2sum问题
func threeSum(nums []int) [][]int {
    var res [][]int
    n := len(nums)
    sort.Ints(nums)
    
    for i := 0; i < n - 2; i ++{
	//注意是nums[i - 1]
        if i == 0 || (i > 0 && nums[i] != nums[i - 1]){
            lo, lh := i + 1, n - 1
            sum := -1 * nums[i]
            for lo < lh{
                if nums[lo] + nums[lh] == sum{
                    res = append(res, []int{nums[i], nums[lo], nums[lh]})
                    fmt.Println([]int{nums[i], nums[lo], nums[lh]})
					//下面两行语句去重
                    for lo < lh && nums[lo] == nums[lo + 1]{lo ++}
                    for lo < lh && nums[lh] == nums[lh - 1]{lh --}
                    lo ++
                    lh --
                } else if nums[lo] + nums[lh] > sum{
                    lh --
                } else {
                	lo ++
                }
            }
        }
    }
    return res
}