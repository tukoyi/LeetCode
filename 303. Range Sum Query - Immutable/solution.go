//参考solution的编写的答案
//

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	var res NumArray
	//res.Sum长度需要是len(nums) + 1，很多类似技巧都是头部尾部做nil处理，让代码更简洁
	res.Sum = make([]int, len(nums) + 1)
	for k, v := range nums{
	res.Sum[k + 1] = res.Sum[k] + v
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	//思路是：分别考虑了sum(i)包括nums[i]和不包括nums[i]的情况，不包括nums[i]才能计算
	return this.Sum[j + 1] - this.Sum[i]
}

/**

Your NumArray object will be instantiated and called as such:
obj := Constructor(nums);
param_1 := obj.SumRange(i,j);
*/