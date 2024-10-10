package problems

// https://leetcode.com/problems/min-stack/description/

type minStackElement struct {
	val, minVal int
}

type MinStack struct {
	elems []minStackElement
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.elems) > 0 {
		minVal = min(minVal, this.GetMin())
	}
	this.elems = append(this.elems, minStackElement{val, minVal})
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1].val
}

func (this *MinStack) GetMin() int {
	return this.elems[len(this.elems)-1].minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
