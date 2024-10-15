package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	nodes := make(map[int]*CloneGraphNode)
	nodes[1] = &CloneGraphNode{Val: 1}
	nodes[2] = &CloneGraphNode{Val: 2}
	nodes[3] = &CloneGraphNode{Val: 3}
	nodes[4] = &CloneGraphNode{Val: 4}
	nodes[1].Neighbors = []*CloneGraphNode{nodes[2], nodes[4]}
	nodes[2].Neighbors = []*CloneGraphNode{nodes[1], nodes[3]}
	nodes[3].Neighbors = []*CloneGraphNode{nodes[2], nodes[4]}
	nodes[4].Neighbors = []*CloneGraphNode{nodes[1], nodes[3]}

	clonedNodes := cloneGraph(nodes[1])
	assert.Equal(t, 1, clonedNodes.Val)
}
