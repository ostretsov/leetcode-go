package problems

import (
	"slices"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/open-the-lock/description/
func openLock(deadends []string, target string) int {
	return openLockV2Faster(deadends, target)
}

func openLockV1Slow(deadends []string, target string) int {
	if slices.Contains(deadends, "0000") {
		return -1
	}
	if target == "0000" {
		return 0
	}

	type move struct {
		state string
		depth int
	}
	queue := kit.SimpleQueue[move]{}
	queue.Enqueue(move{"0000", 0})
	visitedStates := make(map[string]struct{})
	for !queue.Empty() {
		m, _ := queue.Dequeue()
		visitedStates[m.state] = struct{}{}
		for _, nextMove := range allNextMoves(m.state) {
			if nextMove == target {
				return m.depth + 1
			}
			if slices.Contains(deadends, nextMove) {
				continue
			}
			if _, ok := visitedStates[nextMove]; ok {
				continue
			}
			queue.Enqueue(move{nextMove, m.depth + 1})
		}
	}
	return -1
}

var slots = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func allNextMoves(state string) []string {
	rnState := []rune(state)
	var nextMoves []string
	for i := 0; i < len(state); i++ {
		iv := rnState[i]
		ci := rnState[i] - '0'

		nextSlot := slots[(ci+1)%10]
		rnState[i] = nextSlot
		nextMoves = append(nextMoves, string(rnState))

		prevSlot := slots[(ci+9)%10]
		rnState[i] = prevSlot
		nextMoves = append(nextMoves, string(rnState))

		rnState[i] = iv
	}
	return nextMoves
}

func openLockV2Faster(deadends []string, target string) int {
	if slices.Contains(deadends, "0000") {
		return -1
	}
	if target == "0000" {
		return 0
	}

	type move struct {
		state string
		depth int
	}
	queue := kit.SimpleQueue[move]{}
	queue.Enqueue(move{"0000", 0})
	visitedStates := make(map[string]struct{})
	for !queue.Empty() {
		m, _ := queue.Dequeue()
		for _, nextMove := range allNextMoves(m.state) {
			if _, ok := visitedStates[nextMove]; ok {
				continue
			}
			if slices.Contains(deadends, nextMove) {
				continue
			}
			if nextMove == target {
				return m.depth + 1
			}
			visitedStates[nextMove] = struct{}{}
			queue.Enqueue(move{nextMove, m.depth + 1})
		}
	}
	return -1
}
