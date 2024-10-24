package _42_01_matrix

import (
	"leetcode-go/kit"
)

// https://leetcode.com/problems/01-matrix/description/
func updateMatrix(mat [][]int) [][]int {
	return updateMatrixV1Slow(mat)
}

func updateMatrixV1Slow(mat [][]int) [][]int {
	mLen := len(mat)
	nLen := len(mat[0])

	type cell struct {
		m, n, dstZero int
	}

	visitedCells := make(map[int]bool)

	var zeroCells []cell
	for mi := 0; mi < mLen; mi++ {
		for ni := 0; ni < nLen; ni++ {
			if mat[mi][ni] == 0 {
				zeroCells = append(zeroCells, cell{mi, ni, 0})
			}
		}
	}

	for _, zc := range zeroCells {
		queue := kit.SimpleQueue[cell]{}
		visit := func(nextM, nextN, dstZero int) {
			if nextM >= 0 && nextM < mLen && nextN >= 0 && nextN < nLen {
				if mat[nextM][nextN] == 0 {
					return
				}
				if !visitedCells[nextM*10_000+nextN] || mat[nextM][nextN] > dstZero+1 {
					queue.Enqueue(cell{nextM, nextN, dstZero + 1})
				}
			}
		}
		queue.Enqueue(cell{zc.m, zc.n, 0})
		for !queue.Empty() {
			c, _ := queue.Dequeue()
			visitedCells[c.m*10_000+c.n] = true
			mat[c.m][c.n] = c.dstZero

			visit(c.m, c.n+1, c.dstZero)
			visit(c.m, c.n-1, c.dstZero)
			visit(c.m+1, c.n, c.dstZero)
			visit(c.m-1, c.n, c.dstZero)
		}
	}

	return mat
}
