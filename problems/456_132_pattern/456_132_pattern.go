package _56_132_pattern

import (
	"math"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/132-pattern/
func find132pattern(nums []int) bool {
	return monoStack(nums)
}

func bruteForceSlow(nums []int) bool {
	minLeft := math.MaxInt
	for i, n := range nums {
		if i > 0 && nums[i-1] < minLeft {
			minLeft = nums[i-1]
		}
		if n <= minLeft || i == 0 || i == len(nums)-1 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > minLeft && nums[j] < n {
				return true
			}
		}
	}

	return false
}

func monoStack(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	minLeft := make([]int, len(nums))
	ml := math.MaxInt
	minLeft[0] = ml
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < ml {
			ml = nums[i-1]
		}
		minLeft[i] = ml
	}

	stack := kit.SimpleStack[int]{}
	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.Empty() && nums[stack.MustTop()] < nums[i] {
			lessIdx, _ := stack.Pop()
			if nums[lessIdx] > minLeft[i] {
				return true
			}
		}

		stack.Push(i)
	}

	return false
}
