package _96_next_grater_elem1

import "leetcode-go/kit"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make([]int, len(nums2))
	for i := range nextGreater {
		nextGreater[i] = -1
	}
	stack := kit.SimpleStack[int]{}
	for i := 0; i < len(nums2); i++ {
		for !stack.Empty() && nums2[stack.MustTop()] < nums2[i] {
			nextGreater[stack.MustPop()] = nums2[i]
		}
		stack.Push(i)
	}
	set := make(map[int]int) // val => next greater val
	for i, n := range nums2 {
		set[n] = nextGreater[i]
	}
	var result []int
	for _, nq := range nums1 {
		result = append(result, set[nq])
	}
	return result
}
