func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode)
    head.Next = new(ListNode)
    ln := head.Next
    carry := 0
    var x, y int
    
    for l1 != nil || l2 != nil || carry > 0{
        if l1 == nil{
            x = 0
        } else {x = l1.Val}
        if l2 == nil{
            y = 0
        } else {y = l2.Val}
        
        sum := x + y + carry
        carry = sum / 10
        ln.Val = sum % 10

//必须先forward 然后再判断是否生成新ListNode
        if l1 != nil{l1 = l1.Next}
        if l2 != nil{l2 = l2.Next}

        if l1 != nil || l2 != nil || carry > 0{
        	ln.Next = new(ListNode)
        	ln = ln.Next
        }        
         
        
    }
    return head.Next
    
}