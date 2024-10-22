package _33_flood_fill

import (
	"leetcode-go/kit"
)

type pixel struct {
	r, c int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	fromColor := image[sr][sc]
	rLen, cLen := len(image), len(image[0])
	queue := kit.SimpleQueue[pixel]{}
	visited := make(map[int]struct{}) // key = row * 100 + col

	visit := func(r, c int) {
		if r < 0 || r >= rLen || c < 0 || c >= cLen {
			return
		}
		if _, ok := visited[r*100+c]; ok {
			return
		}
		if image[r][c] != fromColor {
			return
		}
		queue.Enqueue(pixel{r, c})
	}

	queue.Enqueue(pixel{sr, sc})
	for !queue.Empty() {
		p, _ := queue.Dequeue()
		image[p.r][p.c] = color
		visited[p.r*100+p.c] = struct{}{}

		visit(p.r+1, p.c)
		visit(p.r-1, p.c)
		visit(p.r, p.c+1)
		visit(p.r, p.c-1)
	}
	return image
}
