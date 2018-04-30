func maxSubArray(nums []int) int {
	
	N := len(nums)
	
	var maxSubArrary int = nums[0] 
	list := make([]int, N)
	list[0] = nums[0]
    
    
	for i := 1; i < N; i ++ {
		//结果数组list只保存maxSubArrary
		//maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i]才是状态转移方程的最终形式
		if list[i - 1] < 0 {
			list[i] = nums[i]
		} else {
			list[i] = nums[i] + list[i - 1]
		}

		if list[i] > maxSubArrary{
			maxSubArrary = list[i]
		}

	} 
    return maxSubArrary
}