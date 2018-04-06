func twoSum(nums []int, target int) []int {
    var result []int
    var k int
    for k = 0; k < len(nums) - 1; k++{
        for i, j := range nums[k + 1:] {
            if nums[k] + j == target
            result = append(result, k, i)
            return  result
        }
    }
    
}