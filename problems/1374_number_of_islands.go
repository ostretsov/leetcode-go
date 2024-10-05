package problems

import (
	"leetcode-go/kit"
	"leetcode-go/utils"
)

func numIslands(grid [][]byte) int {
	mLen := len(grid)
	nLen := len(grid[0])
	visitedLands := utils.ZeroSliceMxN[bool](mLen, nLen)

	type coord struct {
		m, n int
	}

	findFirstNotVisitedLandCoords := func() (int, int, bool) {
		for m := 0; m < mLen; m++ {
			for n := 0; n < nLen; n++ {
				if visitedLands[m][n] {
					continue
				}
				visitedLands[m][n] = true
				if grid[m][n] == '0' {
					continue
				}
				return m, n, true
			}
		}
		return 0, 0, false
	}

	queue := kit.NewSimpleQueue[coord]()
	var cnt int
	for {
		m, n, found := findFirstNotVisitedLandCoords()
		if !found {
			break
		}

		queue.Enqueue(coord{m, n})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			if c.n+1 < nLen && !visitedLands[c.m][c.n+1] {
				visitedLands[c.m][c.n+1] = true
				if grid[c.m][c.n+1] == '1' {
					queue.Enqueue(coord{c.m, c.n + 1})
				}
			}
			if c.m+1 < mLen && !visitedLands[c.m+1][c.n] {
				visitedLands[c.m+1][c.n] = true
				if grid[c.m+1][c.n] == '1' {
					queue.Enqueue(coord{c.m + 1, c.n})
				}
			}
			if c.n-1 >= 0 && !visitedLands[c.m][c.n-1] {
				visitedLands[c.m][c.n-1] = true
				if grid[c.m][c.n-1] == '1' {
					queue.Enqueue(coord{c.m, c.n - 1})
				}
			}
			if c.m-1 >= 0 && !visitedLands[c.m-1][c.n] {
				visitedLands[c.m-1][c.n] = true
				if grid[c.m-1][c.n] == '1' {
					queue.Enqueue(coord{c.m - 1, c.n})
				}
			}
		}
		cnt++
	}
	return cnt
}
