package _03_next_gt_elem2

import (
	"leetcode-go/kit"
)

// https://leetcode.com/problems/next-greater-element-ii/
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = -1
	}

	decStack := kit.SimpleStack[int]{}
	for k := 0; k < len(nums)*2-1; k++ {
		i := k % len(nums)
		for !decStack.Empty() {
			idx, _ := decStack.Top()
			if nums[idx] >= nums[i] { // 3 2 3?
				break
			}
			for !decStack.Empty() {
				topIdx, _ := decStack.Top()
				if nums[topIdx] >= nums[i] {
					break
				}
				idx, _ = decStack.Pop()
				if res[idx] == -1 {
					res[idx] = nums[i]
				}
			}
		}
		decStack.Push(i)
	}
	return res
}
