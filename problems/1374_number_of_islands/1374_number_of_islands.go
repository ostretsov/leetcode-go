package _374_number_of_islands

import (
	"leetcode-go/kit"
	"leetcode-go/utils"
)

func numIslands(grid [][]byte) int {
	return dfsNumIslands(grid)
}

func dfsNumIslands(grid [][]byte) int {
	mLen := len(grid)
	nLen := len(grid[0])

	var numIslands int
	for m := 0; m < mLen; m++ {
		for n := 0; n < nLen; n++ {
			if grid[m][n] == 'v' || grid[m][n] == '0' {
				continue
			}
			dfsNumIslandsVisitLoop(grid, mLen, nLen, m, n)
			numIslands++
		}
	}
	return numIslands
}

func dfsNumIslandsVisitLoop(grid [][]byte, mLen, nLen, m, n int) {
	if m < 0 || m >= mLen || n < 0 || n >= nLen {
		return
	}
	if grid[m][n] == 'v' || grid[m][n] == '0' {
		return
	}
	grid[m][n] = 'v'
	dfsNumIslandsVisitLoop(grid, mLen, nLen, m+1, n)
	dfsNumIslandsVisitLoop(grid, mLen, nLen, m-1, n)
	dfsNumIslandsVisitLoop(grid, mLen, nLen, m, n+1)
	dfsNumIslandsVisitLoop(grid, mLen, nLen, m, n-1)
}

func bfsNumIslandsV2(grid [][]byte) int {
	mLen := len(grid)
	nLen := len(grid[0])
	visitedLands := utils.ZeroSliceMxN[bool](mLen, nLen)

	var mSearchIter, nSearchIter int
	findFirstNotVisitedLandCoords := func() (int, int, bool) {
		for ; mSearchIter < mLen; mSearchIter++ {
			for ; nSearchIter < nLen; nSearchIter++ {
				if visitedLands[mSearchIter][nSearchIter] {
					continue
				}
				visitedLands[mSearchIter][nSearchIter] = true
				if grid[mSearchIter][nSearchIter] == '0' {
					continue
				}
				return mSearchIter, nSearchIter, true
			}
			nSearchIter = 0
		}
		return 0, 0, false
	}

	type coords struct {
		m, n int
	}
	queue := kit.NewSimpleQueue[coords]()
	enqueueIfNotVisitedLand := func(m, n int) {
		if n < 0 || n >= nLen || m < 0 || m >= mLen {
			return
		}
		if visitedLands[m][n] {
			return
		}
		visitedLands[m][n] = true
		if grid[m][n] == '1' {
			queue.Enqueue(coords{m, n})
		}
	}
	var cnt int
	for {
		m, n, found := findFirstNotVisitedLandCoords()
		if !found {
			break
		}

		queue.Enqueue(coords{m, n})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			enqueueIfNotVisitedLand(c.m, c.n+1)
			enqueueIfNotVisitedLand(c.m+1, c.n)
			enqueueIfNotVisitedLand(c.m, c.n-1)
			enqueueIfNotVisitedLand(c.m-1, c.n)
		}
		cnt++
	}
	return cnt
}

func bfsNumIslandsV3(grid [][]byte) int {
	mLen := len(grid)
	nLen := len(grid[0])

	var mSearchIter, nSearchIter int
	findFirstNotVisitedLandCoords := func() (int, int, bool) {
		for ; mSearchIter < mLen; mSearchIter++ {
			for ; nSearchIter < nLen; nSearchIter++ {
				if grid[mSearchIter][nSearchIter] == 'v' {
					continue
				}
				if grid[mSearchIter][nSearchIter] == '0' {
					grid[mSearchIter][nSearchIter] = 'v'
					continue
				}
				grid[mSearchIter][nSearchIter] = 'v'
				return mSearchIter, nSearchIter, true
			}
			nSearchIter = 0
		}
		return 0, 0, false
	}

	type coords struct {
		m, n int
	}
	queue := kit.NewSimpleQueue[coords]()
	enqueueIfNotVisitedLand := func(m, n int) {
		if n < 0 || n >= nLen || m < 0 || m >= mLen {
			return
		}
		if grid[m][n] == 'v' {
			return
		}
		if grid[m][n] == '1' {
			queue.Enqueue(coords{m, n})
		}
		grid[m][n] = 'v'
	}
	var cnt int
	for {
		m, n, found := findFirstNotVisitedLandCoords()
		if !found {
			break
		}

		queue.Enqueue(coords{m, n})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			enqueueIfNotVisitedLand(c.m, c.n+1)
			enqueueIfNotVisitedLand(c.m+1, c.n)
			enqueueIfNotVisitedLand(c.m, c.n-1)
			enqueueIfNotVisitedLand(c.m-1, c.n)
		}
		cnt++
	}
	return cnt
}
