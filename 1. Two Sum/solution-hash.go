func twoSum(nums []int, target int) []int {
	var result []int
	var hstab = make(map[int]int)
	var v int
	for k, v := range nums{
		hstab[v] = k
	}
    
	for i, j := range nums{
		v = target - j
		mapv, ok := hstab[v]
		if ok && i != mapv {
		    result = append(result, i, hstab[v])
            return  result	
		}
	}
	return  result	

}