package _39_daily_temperatures

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailyTemperatures(t *testing.T) {
	got := dailyTemperatures([]int{72, 76, 73})
	assert.Equal(t, []int{1, 0, 0}, got)

	got = dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70})
	assert.Equal(t, []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}, got)

	got = dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, got)

	got = dailyTemperatures([]int{30, 40, 50, 60})
	assert.Equal(t, []int{1, 1, 1, 0}, got)

	got = dailyTemperatures([]int{30, 60, 90})
	assert.Equal(t, []int{1, 1, 0}, got)

	input := slices.Repeat([]int{99}, 99_999)
	input = append(input, 100)
	dailyTemperatures(input)
}
