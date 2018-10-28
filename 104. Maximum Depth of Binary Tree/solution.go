/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 /* 递归算法，最大深度为max（leftHeight,rightHeight）+ 1
	
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } else {
        leftHeight := maxDepth(root.Left)
        rightHeight := maxDepth(root.Right)
        
        if leftHeight > rightHeight {
            return leftHeight + 1
        } else {
            return rightHeight + 1
        }
    }
}