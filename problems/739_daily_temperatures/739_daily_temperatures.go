package _39_daily_temperatures

import "leetcode-go/kit"

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures(temperatures []int) []int {
	return dailyTemperaturesV2Stack(temperatures)
}

func dailyTemperaturesV1Naive(temperatures []int) []int {
	temps := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		j := i + 1
		var waitForWarmerTempDays int
		for ; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				waitForWarmerTempDays = j - i
				break
			}
		}
		temps[i] = waitForWarmerTempDays
	}
	return temps
}

func dailyTemperaturesV2Stack(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := kit.SimpleStack[int]{}
	for ti, tempToday := range temperatures {
		for !stack.Empty() {
			recentTempInd, _ := stack.Top()
			if tempToday <= temperatures[recentTempInd] {
				break
			}
			stack.Pop()
			results[recentTempInd] = ti - recentTempInd
		}
		stack.Push(ti)
	}
	return results
}
