package _32_queue_using_stacks

import "leetcode-go/kit"

// https://leetcode.com/problems/implement-queue-using-stacks/description/
type MyQueue struct {
	stack kit.SimpleStack[int]
}

func Constructor() MyQueue {
	return MyQueue{
		stack: kit.SimpleStack[int]{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack.Push(x)
}

func (this *MyQueue) Pop() int {
	tmpStack := kit.SimpleStack[int]{}
	for !this.stack.Empty() {
		n, _ := this.stack.Pop()
		tmpStack.Push(n)
	}
	queueHead, _ := tmpStack.Pop()
	for !tmpStack.Empty() {
		n, _ := tmpStack.Pop()
		this.stack.Push(n)
	}
	return queueHead
}

func (this *MyQueue) Peek() int {
	tmpStack := kit.SimpleStack[int]{}
	for !this.stack.Empty() {
		n, _ := this.stack.Pop()
		tmpStack.Push(n)
	}
	queueHead, _ := tmpStack.Top()
	for !tmpStack.Empty() {
		n, _ := tmpStack.Pop()
		this.stack.Push(n)
	}
	return queueHead
}

func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
