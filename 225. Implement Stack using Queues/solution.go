type MyStack struct {
    Stack   []int
    Length  int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    ret := MyStack{
	//这里list的容量要是0， 否则影响pop输出结果
        Stack:  make([]int, 0),	
        Length: 0,
    }
    return ret
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Stack = append(this.Stack, x)
    this.Length ++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    
    ret := this.Stack[this.Length - 1]
    this.Stack = this.Stack[:this.Length - 1]
    this.Length --
    return ret
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.Stack[this.Length - 1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if this.Length == 0 {
        return true
    } else {
        return false
    }
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */