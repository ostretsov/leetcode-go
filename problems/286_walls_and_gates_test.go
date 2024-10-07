package problems

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_wallsAndGates(t *testing.T) {
	const (
		r = 2147483647
		G = 0
		W = -1
	)

	map1 := [][]int{
		{r, W, G, r},
		{r, r, r, W},
		{r, W, r, W},
		{G, W, r, r},
	}
	wallsAndGates(map1)
	assert.Equal(t, []int{3, -1, 0, 1}, map1[0])
	assert.Equal(t, []int{2, 2, 1, -1}, map1[1])
	assert.Equal(t, []int{1, -1, 2, -1}, map1[2])
	assert.Equal(t, []int{0, -1, 3, 4}, map1[3])

	map2 := [][]int{
		{G, r, W, r, W, W, G, W, r, r, r, G, G, r, W, G, r, W, r, r, G, G, W, G, r, W, r, W, W, W, r, r, G, r, r, W, G, G, G, G, r, r, G, r, r, G},
		{W, G, r, W, W, r, r, r, G, G, G, r, G, W, G, W, G, W, r, G, W, W, r, W, r, G, r, r, W, W, r, G, r, r, G, G, G, W, r, W, W, r, G, W, G, G},
		{W, G, G, r, G, r, G, r, G, G, G, G, W, r, r, W, r, W, r, W, r, W, r, W, r, G, G, r, W, G, W, G, G, G, G, W, G, W, G, W, r, r, W, W, r, G},
		{G, W, r, r, G, r, G, r, W, r, W, r, r, r, r, W, G, r, G, W, W, W, G, G, G, W, r, G, W, r, r, W, W, G, W, r, G, r, G, G, r, W, W, r, r, G},
		{W, W, G, G, r, r, G, G, r, W, G, W, G, G, r, G, r, G, r, G, r, r, G, r, W, G, G, r, G, r, r, r, W, r, W, r, G, G, W, r, W, G, W, r, W, r},
		{G, W, r, W, G, W, W, r, r, G, r, G, W, G, r, W, r, r, G, W, G, G, r, W, G, W, r, W, r, G, G, G, G, r, W, G, r, r, W, r, G, W, r, G, r, r},
		{G, G, G, r, W, G, r, W, r, W, W, W, W, W, W, W, W, G, G, r, W, G, r, G, r, W, G, W, W, r, G, r, W, G, G, W, G, W, r, W, r, r, W, G, G, G},
		{r, r, r, G, G, r, G, W, G, W, W, W, G, r, W, W, W, r, W, r, W, G, r, W, r, G, r, W, W, r, W, G, G, r, W, r, G, G, G, r, G, W, W, G, G, W},
		{W, G, W, r, r, G, r, r, W, G, W, r, G, G, W, G, W, W, G, G, W, W, r, r, r, G, r, W, W, G, r, W, G, r, W, W, G, r, W, r, r, G, r, W, r, r},
		{W, r, W, r, W, r, W, W, W, G, W, G, r, G, G, G, W, W, W, G, G, W, W, W, G, W, r, r, W, G, r, r, r, G, r, W, r, G, W, W, r, W, r, G, r, W},
		{G, G, r, G, r, r, G, W, W, W, G, W, W, W, G, G, W, r, r, G, G, G, W, G, W, W, G, W, W, G, G, G, G, r, W, r, r, W, r, W, r, W, G, G, W, W},
		{W, r, r, G, G, r, W, r, G, r, G, G, r, r, G, W, G, W, r, r, W, r, W, W, r, r, G, W, G, r, G, G, r, r, W, r, G, W, G, G, W, G, G, G, W, W},
		{G, W, r, G, W, W, G, W, r, W, W, W, W, r, W, r, r, r, r, W, W, r, W, r, W, r, W, r, W, r, G, r, G, W, W, G, G, W, G, r, W, W, G, r, G, G},
		{W, W, G, W, r, W, G, W, r, r, r, W, G, W, W, G, G, G, G, G, G, r, G, W, r, G, G, W, r, r, r, r, W, G, G, G, G, W, r, G, G, G, G, W, r, W},
		{r, G, G, r, G, r, W, G, G, r, r, r, r, r, G, W, G, W, W, r, W, G, W, W, G, r, r, W, W, r, r, W, r, r, W, r, G, G, G, W, W, G, G, G, G, r},
		{G, W, W, G, G, W, W, W, G, G, r, r, r, r, W, W, G, W, W, W, G, r, r, G, r, G, G, G, G, r, G, r, r, G, W, r, W, W, r, r, r, W, r, G, W, r},
		{W, W, W, W, r, G, G, r, G, G, r, G, W, G, r, G, G, r, G, W, G, G, G, G, G, W, G, r, r, G, G, r, G, G, r, r, W, r, r, r, G, W, W, G, r, r},
		{r, W, W, G, r, W, G, G, W, W, r, G, G, W, G, r, r, W, r, W, G, r, W, W, G, r, W, G, W, W, W, W, r, W, W, G, W, r, W, r, W, r, G, W, G, G},
		{r, r, W, r, W, G, G, G, r, r, G, W, W, G, W, r, r, W, G, r, r, r, G, G, r, r, G, r, W, r, r, G, r, G, r, r, r, G, W, r, r, r, r, W, W, r},
	}
	start := time.Now()
	wallsAndGates(map2)
	execTime := time.Since(start)
	assert.Less(t, execTime, 1*time.Second)

	map3 := [][]int{
		{G, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, G, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, G, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, G, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, G, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, G, r, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, G, r, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, G, r, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, G, r, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, G, r, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, G, r, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, G, r, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, G, r, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, G, r, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, G, r, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, G, r, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, G, r, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, G, r, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, G, r},
		{r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, G},
	}
	wallsAndGates(map3)
}
