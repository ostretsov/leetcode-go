package problems

// Rename to "Node"
type CloneGraphNode struct {
	Val       int
	Neighbors []*CloneGraphNode
}

// https://leetcode.com/problems/clone-graph/description/
func cloneGraph(node *CloneGraphNode) *CloneGraphNode {
	if node == nil {
		return nil
	}
	clonedNode := &CloneGraphNode{}
	visited := make(map[int]*CloneGraphNode)
	cloneGraphRecursion(node, clonedNode, visited)
	return clonedNode
}

func cloneGraphRecursion(origNode, clonedNode *CloneGraphNode, visited map[int]*CloneGraphNode) {
	clonedNode.Val = origNode.Val
	visited[origNode.Val] = clonedNode
	for _, neighbor := range origNode.Neighbors {
		clonedNeighbor, ok := visited[neighbor.Val]
		if !ok {
			clonedNeighbor = &CloneGraphNode{}
			cloneGraphRecursion(neighbor, clonedNeighbor, visited)
		}
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
	}
}
