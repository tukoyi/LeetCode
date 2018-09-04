//保存prev 和next节点信息

package main

func reverseList(head *ListNode) *ListNode {
	var cur, prev, next *ListNode

	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}
