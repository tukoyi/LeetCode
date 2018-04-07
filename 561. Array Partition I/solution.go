//时间复杂度O(2n)
//beat 20%
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    sum := 0
    for k, v := range nums{
        if k % 2 == 0{
          sum += v  
        }       
    }
    return sum
}

//稍微改一下beat 93%
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    sum := 0
    for k :=0; k < len(nums); {
          sum += nums[k]
          k = k + 2
    }
    return sum
}