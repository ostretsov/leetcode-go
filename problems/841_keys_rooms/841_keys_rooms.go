package _41_keys_rooms

import "leetcode-go/kit"

// https://leetcode.com/problems/keys-and-rooms/
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	queue := kit.SimpleQueue[int]{}
	queue.Enqueue(0)
	for !queue.Empty() {
		ri, _ := queue.Dequeue()
		for _, rkey := range rooms[ri] {
			if !visited[rkey] {
				queue.Enqueue(rkey)
			}
		}
		visited[ri] = true
	}

	for _, vr := range visited {
		if !vr {
			return false
		}
	}
	return true
}
