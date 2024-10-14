package utils

import "math"

func FloatEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
