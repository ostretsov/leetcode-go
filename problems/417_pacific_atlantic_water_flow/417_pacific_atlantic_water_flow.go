package _17_pacific_atlantic_water_flow

import (
	"leetcode-go/kit"
	"leetcode-go/utils"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/
func pacificAtlantic(heights [][]int) [][]int {
	mLen, nLen := len(heights), len(heights[0])
	const atlantic = 1
	const pacific = 2
	const atlanticAndPacific = 3
	visited := utils.ZeroSliceMxN[int](mLen, nLen)

	type cell struct{ m, n, ocean int }
	queue := kit.SimpleQueue[cell]{}

	visit := func(m, n, prevHeight, ocean int) {
		if m < 0 || m >= mLen || n < 0 || n >= nLen {
			return
		}
		if visited[m][n]&ocean > 0 { // visited by the ocean
			return
		}
		if prevHeight > heights[m][n] {
			return
		}
		queue.Enqueue(cell{m, n, ocean})
	}

	for mi := 0; mi < mLen; mi++ {
		queue.Enqueue(cell{mi, 0, pacific})
		queue.Enqueue(cell{mi, nLen - 1, atlantic})
	}
	for ni := 0; ni < nLen; ni++ {
		if ni > 0 {
			queue.Enqueue(cell{0, ni, pacific})
		}
		if ni < nLen-1 {
			queue.Enqueue(cell{mLen - 1, ni, atlantic})
		}
	}

	for !queue.Empty() {
		c, _ := queue.Dequeue()
		if visited[c.m][c.n]&c.ocean > 0 {
			continue
		}
		visited[c.m][c.n] |= c.ocean

		currentHeight := heights[c.m][c.n]
		visit(c.m, c.n+1, currentHeight, c.ocean)
		visit(c.m, c.n-1, currentHeight, c.ocean)
		visit(c.m-1, c.n, currentHeight, c.ocean)
		visit(c.m+1, c.n, currentHeight, c.ocean)
	}

	var res [][]int
	for mi := 0; mi < mLen; mi++ {
		for ni := 0; ni < nLen; ni++ {
			if visited[mi][ni] == atlanticAndPacific {
				res = append(res, []int{mi, ni})
			}
		}
	}
	return res
}
