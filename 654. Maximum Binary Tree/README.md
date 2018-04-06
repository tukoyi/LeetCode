1、错误的递归方式：
    tnode.Val = nums[maxk]	
    tnode.Left = constructMaximumBinaryTree(nums[:maxk])
    tnode.Right = constructMaximumBinaryTree(nums[maxk + 1:])		
在num=[11, 22]的情况下，tnode.Left = constructMaximumBinaryTree(nums[2:])
out of range产生报错

2、必须用指证
    if len(nums) == 0{   
        return nil
	} 