package main

//本题有两种思路，一种是list 翻转，一种是使用粘的特性，这里使用了第二种方方法，特地实现了栈
//栈的实现可以参考 https://github.com/xcltapestry/xclpkg/blob/master/algorithm/stack.go
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	Elements []int
}

func (stack *Stack) Push(data int) {
	stack.Elements = append(stack.Elements, data)
}

func (stack *Stack) Pop() int {
	if len(stack.Elements) > 0 {
		res := stack.Elements[len(stack.Elements)-1]
		stack.Elements = stack.Elements[:len(stack.Elements)-1]
		return res
	}
	return 0
}

func (stack *Stack) Empty() bool {
	if len(stack.Elements) == 0 || stack.Elements == nil {
		return true
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 Stack
	list1, list2 := l1, l2

	//create stack
	for list1 != nil {
		s1.Push(list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		s2.Push(list2.Val)
		list2 = list2.Next
	}
	fmt.Printf("s1:%v s2:%v", s1.Elements, s2.Elements)

	var sum int
	var head *ListNode
	for !s1.Empty() || !s2.Empty() {

		if !s1.Empty() {
			sum += s1.Pop()
		}
		if !s2.Empty() {
			sum += s2.Pop()
		}

		var newNode ListNode
		newNode.Val = sum % 10
		newNode.Next = head
		head = &newNode
		sum = sum / 10

	}

	//最高位有进位的情况容易忽略
	if sum > 0 {
		var newNode ListNode
		newNode.Val = sum
		newNode.Next = head
		head = &newNode
	}

	return head
}

func main() {
	l1 := ListNode{7, &ListNode{2, &ListNode{4, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l3 := addTwoNumbers(&l1, &l2)

	for l3 != nil {
		fmt.Printf("->%d", l3.Val)
		l3 = l3.Next
	}
}
