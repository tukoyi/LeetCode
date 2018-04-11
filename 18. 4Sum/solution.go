func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var res [][]int
    nlimit := len(nums)
    var lo,lh int
    fmt.Println(nums)
    if nlimit < 4{
    	return res
    }

    for i := 0; i < nlimit - 3; i++{
        if i == 0 || (i > 0 && nums[i] != nums[i - 1]){
	        for j := i + 1; j < nlimit -2; j++{
	            if j == i + 1 || (j > i + 1 && nums[j] != nums[j - 1]){
	            	negsum2 := target - nums[i] - nums[j]
	            	lo, lh = j + 1, nlimit - 1
	            	fmt.Println(i, j, lo, lh)
		            for lo < lh{
		                if nums[lo] + nums[lh] == negsum2{
		                    res = append(res, []int{nums[i], nums[j], nums[lo], nums[lh]})
		                    //nums[lo] == nums[lo + 1] && lo < lh的话会报index out of range,原因是S := []int{0, 0, 0,0}时，会执行
		                    //nums[3] == nums[3 + 1] && lo < lh,导致报错，因此这种这类问题，先判断index是否合法在判断值是否相等
		                    for lo < lh && nums[lo] == nums[lo + 1]  {lo ++}
		                    for nums[lh] == nums[lh - 1] && lo < lh {lh --}
		                    lh --
		                    lo ++
		                } else if nums[lo] + nums[lh] > negsum2{
		                    lh --
		                } else{
		                    lo ++
		                }

		            }
	            }
	        }
	    }
    }
    return res
    
}