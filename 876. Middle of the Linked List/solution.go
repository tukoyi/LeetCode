//只需要遍历一次list，将midddle的位置保存下来

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var mid, cur *ListNode
    var cnt = 1
    mid = head
    cur = head
    
    for cur != nil {
        if cnt % 2 == 0 {
            mid = mid.Next
        }
        cur = cur.Next
        cnt += 1
    }
    return mid
}