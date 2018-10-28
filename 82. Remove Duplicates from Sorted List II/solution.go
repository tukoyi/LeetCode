/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	newHead := &ListNode{}
	slow := newHead
	fast := head
	slow.Next = fast

	for fast != nil {
		for fast.Next != nil && fast.Val == fast.Next.Val{
			fast = fast.Next
		}

		if slow.Next != fast {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return newHead.Next

}