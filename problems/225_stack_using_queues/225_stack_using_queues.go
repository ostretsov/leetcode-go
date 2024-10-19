package _25_stack_using_queues

import "leetcode-go/kit"

// https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	queue kit.SimpleQueue[int]
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)
}

func (this *MyStack) Pop() int {
	tmpQueue := kit.SimpleQueue[int]{}
	var lastElem int
	initQueueLen := this.queue.Len()
	for i := 0; !this.queue.Empty(); i++ {
		lastElem, _ = this.queue.Dequeue()
		if i < initQueueLen-1 {
			tmpQueue.Enqueue(lastElem)
		}
	}
	this.queue = tmpQueue
	return lastElem
}

func (this *MyStack) Top() int {
	tmpQueue := kit.SimpleQueue[int]{}
	var lastElem int
	for !this.queue.Empty() {
		lastElem, _ = this.queue.Dequeue()
		tmpQueue.Enqueue(lastElem)
	}
	this.queue = tmpQueue
	return lastElem
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
