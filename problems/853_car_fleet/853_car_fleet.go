package _53_car_fleet

import (
	"slices"

	"leetcode-go/kit"
	"leetcode-go/utils"
)

// https://leetcode.com/problems/car-fleet/
func carFleet(target int, position []int, speed []int) int {
	type car struct {
		position int
		ttf      float64 // time to finish
	}
	var cars []car
	for i := range position {
		ttf := float64(target-position[i]) / float64(speed[i])
		cars = append(cars, car{position[i], ttf})
	}
	slices.SortFunc(cars, func(a, b car) int {
		return a.position - b.position
	})
	stack := kit.SimpleStack[car]{}
	stack.Push(cars[0])
	for i := 1; i < len(cars); i++ {
		curCar := cars[i]
		for !stack.Empty() {
			prvCar, _ := stack.Top()
			makeFleet := utils.FloatEqual(curCar.ttf, prvCar.ttf, 1e-9) || curCar.ttf > prvCar.ttf
			if makeFleet {
				stack.Pop()
				continue
			}
			break
		}
		stack.Push(curCar)
	}
	return stack.Len()
}
