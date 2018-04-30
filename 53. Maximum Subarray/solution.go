func maxSubArray(nums []int) int {
	
	N := len(nums)
	//maxSubArrary初始值不能为0，因为元素值可能为负数，nums只有一个元素的情况下会error
	//状态转移方程Dn = Dn-1 + nums(n), Dn为以nums第N个元素为结尾的subarrary之和
	//时间复杂度O(n2),只打败了8%的对手，时间复杂度需要优化
	
	list := make([]int, N)
	maxSubArrary = nums[0]
	
    
    
	for i := 0; i < N; i ++ {
		for j := 0; j <= i; j ++ {
			list[j] = list[j] + nums[i]
			if list[j] > maxSubArrary{
				maxSubArrary = list[j]
			}
			// fmt.Println(list)
		}

	} 
    return maxSubArrary
}