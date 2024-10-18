package _79_perfect_squares

import "leetcode-go/kit"

// https://leetcode.com/problems/perfect-squares/description/
func numSquares(n int) int {
	return numSquaresV2Faster(n)
}

func numSquaresV1Slow(n int) int {
	var sqSet []int
	for i := 1; ; i++ {
		sq := i * i
		if sq > n {
			break
		}
		sqSet = append(sqSet, sq)
	}
	type round struct{ sum, depth int }
	queue := kit.SimpleQueue[round]{}
	queue.Enqueue(round{0, 0})
	for !queue.Empty() {
		r, _ := queue.Dequeue()
		for _, sq := range sqSet {
			sum := r.sum + sq
			if sum == n {
				return r.depth + 1
			}
			if sum > n {
				break
			}
			queue.Enqueue(round{sum, r.depth + 1})
		}
	}
	return -1
}

func numSquaresV2Faster(n int) int {
	var sqSet []int
	for i := 1; ; i++ {
		sq := i * i
		if sq > n {
			break
		}
		sqSet = append(sqSet, sq)
	}
	type round struct{ sum, depth int }
	queue := kit.SimpleQueue[round]{}
	queue.Enqueue(round{0, 0})
	reachedSums := make(map[int]struct{})
	for !queue.Empty() {
		r, _ := queue.Dequeue()
		for _, sq := range sqSet {
			sum := r.sum + sq
			if sum > n {
				break
			}
			if _, ok := reachedSums[sum]; ok {
				continue
			}
			if sum == n {
				return r.depth + 1
			}
			reachedSums[sum] = struct{}{}
			queue.Enqueue(round{sum, r.depth + 1})
		}
	}
	return -1
}
