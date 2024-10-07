package problems

import (
	"leetcode-go/kit"
	"leetcode-go/utils"
)

// https://leetcode.com/problems/walls-and-gates/description/
func wallsAndGates(rooms [][]int) {
	wallsAndGatesV4Short(rooms)
}

func wallsAndGatesV1Slow(rooms [][]int) {
	mLen := len(rooms)
	nLen := len(rooms[0])

	gSearchM, gSearchN := 0, 0
	nextGates := func() (int, int, bool) {
		for ; gSearchM < mLen; gSearchM++ {
			for ; gSearchN < nLen; gSearchN++ {
				if rooms[gSearchM][gSearchN] == 0 {
					defer func() { gSearchN++ }()
					return gSearchM, gSearchN, true
				}
			}
			gSearchN = 0
		}
		return 0, 0, false
	}

	type coords struct {
		m, n      int
		dstFromGt int
	}

	queue := kit.SimpleQueue[coords]{}
	visitNotVisitedRoom := func(visitedRooms [][]bool, m, n, dstFromGt int) {
		if m < 0 || m >= mLen || n < 0 || n >= nLen {
			return
		}
		if rooms[m][n] == -1 || visitedRooms[m][n] {
			return
		}
		queue.Enqueue(coords{m, n, dstFromGt})
	}

	for {
		mGt, nGt, found := nextGates()
		if !found {
			break
		}

		visitedRooms := utils.ZeroSliceMxN[bool](mLen, nLen)
		queue.Enqueue(coords{mGt, nGt, 0})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			if c.dstFromGt < rooms[c.m][c.n] {
				rooms[c.m][c.n] = c.dstFromGt
			}
			visitedRooms[c.m][c.n] = true

			visitNotVisitedRoom(visitedRooms, c.m, c.n+1, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m, c.n-1, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m+1, c.n, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m-1, c.n, c.dstFromGt+1)
		}
	}
}

// Memory Limit Exceeded
func wallsAndGatesV2Faster(rooms [][]int) {
	mLen := len(rooms)
	nLen := len(rooms[0])

	gSearchM, gSearchN := 0, 0
	nextGates := func() (int, int, bool) {
		for ; gSearchM < mLen; gSearchM++ {
			for ; gSearchN < nLen; gSearchN++ {
				if rooms[gSearchM][gSearchN] == 0 {
					defer func() { gSearchN++ }()
					return gSearchM, gSearchN, true
				}
			}
			gSearchN = 0
		}
		return 0, 0, false
	}

	type coords struct {
		m, n      int
		dstFromGt int
	}

	queue := kit.SimpleQueue[coords]{}
	visitNotVisitedRoom := func(visitedRooms [][]bool, m, n, dstFromGt int) {
		if m < 0 || m >= mLen || n < 0 || n >= nLen {
			return
		}
		if rooms[m][n] == -1 || rooms[m][n] == 0 || visitedRooms[m][n] {
			return
		}
		queue.Enqueue(coords{m, n, dstFromGt})
	}

	for {
		mGt, nGt, found := nextGates()
		if !found {
			break
		}

		visitedRooms := utils.ZeroSliceMxN[bool](mLen, nLen)
		queue.Enqueue(coords{mGt, nGt, 0})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			if c.dstFromGt < rooms[c.m][c.n] {
				rooms[c.m][c.n] = c.dstFromGt
			}
			visitedRooms[c.m][c.n] = true

			visitNotVisitedRoom(visitedRooms, c.m, c.n+1, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m, c.n-1, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m+1, c.n, c.dstFromGt+1)
			visitNotVisitedRoom(visitedRooms, c.m-1, c.n, c.dstFromGt+1)
		}
	}
}

func wallsAndGatesV3MemEfficient(rooms [][]int) {
	mLen := len(rooms)
	nLen := len(rooms[0])

	gSearchM, gSearchN := 0, 0
	nextGates := func() (int, int, bool) {
		for ; gSearchM < mLen; gSearchM++ {
			for ; gSearchN < nLen; gSearchN++ {
				if rooms[gSearchM][gSearchN] == 0 {
					defer func() { gSearchN++ }()
					return gSearchM, gSearchN, true
				}
			}
			gSearchN = 0
		}
		return 0, 0, false
	}

	type coords struct {
		m, n      int
		dstFromGt int
	}

	for {
		mGt, nGt, found := nextGates()
		if !found {
			break
		}

		queue := kit.SimpleQueue[coords]{}
		visitNotVisitedRoom := func(m, n, dstFromGt int) {
			if m < 0 || m >= mLen || n < 0 || n >= nLen {
				return
			}
			if rooms[m][n] == -1 || rooms[m][n] == 0 || rooms[m][n] <= dstFromGt {
				return
			}
			rooms[m][n] = dstFromGt
			queue.Enqueue(coords{m, n, dstFromGt})
		}

		queue.Enqueue(coords{mGt, nGt, 0})
		for !queue.Empty() {
			c, _ := queue.Dequeue()

			visitNotVisitedRoom(c.m, c.n+1, c.dstFromGt+1)
			visitNotVisitedRoom(c.m, c.n-1, c.dstFromGt+1)
			visitNotVisitedRoom(c.m+1, c.n, c.dstFromGt+1)
			visitNotVisitedRoom(c.m-1, c.n, c.dstFromGt+1)
		}
	}
}

func wallsAndGatesV4Short(rooms [][]int) {
	mLen := len(rooms)
	nLen := len(rooms[0])

	type coords struct {
		m, n      int
		dstFromGt int
	}
	queue := kit.SimpleQueue[coords]{}
	visitNotVisitedRoom := func(m, n, dstFromGt int) {
		if m < 0 || m >= mLen || n < 0 || n >= nLen {
			return
		}
		if rooms[m][n] == -1 || rooms[m][n] == 0 || rooms[m][n] <= dstFromGt {
			return
		}
		rooms[m][n] = dstFromGt
		queue.Enqueue(coords{m, n, dstFromGt})
	}

	for mi := 0; mi < mLen; mi++ {
		for ni := 0; ni < nLen; ni++ {
			if rooms[mi][ni] == 0 {
				queue.Enqueue(coords{mi, ni, 0})
			}
		}
	}

	for !queue.Empty() {
		c, _ := queue.Dequeue()
		visitNotVisitedRoom(c.m, c.n+1, c.dstFromGt+1)
		visitNotVisitedRoom(c.m, c.n-1, c.dstFromGt+1)
		visitNotVisitedRoom(c.m+1, c.n, c.dstFromGt+1)
		visitNotVisitedRoom(c.m-1, c.n, c.dstFromGt+1)
	}
}
