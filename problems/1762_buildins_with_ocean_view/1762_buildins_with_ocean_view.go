package _762_buildins_with_ocean_view

import (
	"math"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/buildings-with-an-ocean-view/
func findBuildings(heights []int) []int {
	heights = append(heights, math.MaxInt)
	stack := kit.SimpleStack[int]{}
	for i := 0; i < len(heights); i++ {
		for !stack.Empty() {
			if heights[stack.MustTop()] > heights[i] {
				break
			}
			heights[stack.MustPop()] = heights[i]
		}
		stack.Push(i)
	}
	var result []int
	for i := 0; i < len(heights)-1; i++ {
		if heights[i] == math.MaxInt {
			result = append(result, i)
		}
	}
	return result
}
