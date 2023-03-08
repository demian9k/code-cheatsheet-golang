package shortest_path

import (
	"testing"
)

func TestNewSimpleDijkstra(t *testing.T) {

	nodeCount := 6
	edgeCount := 11
	graphRels := []GraphRel{
		{1, 2, 2},
		{1, 3, 5},
		{1, 4, 1},
		{2, 3, 3},
		{2, 4, 2},
		{3, 2, 3},
		{3, 6, 5},
		{4, 3, 3},
		{4, 5, 1},
		{5, 3, 1},
		{5, 6, 2},
	}

	dijkstra := NewSimpleDijkstra(nodeCount, edgeCount, graphRels)
	dijkstra.Execute(1)
	dijkstra.PrintResult()
}

func TestNewFastDijkstra(t *testing.T) {

	nodeCount := 6
	edgeCount := 11
	graphRels := []GraphRel{
		{1, 2, 2},
		{1, 3, 5},
		{1, 4, 1},
		{2, 3, 3},
		{2, 4, 2},
		{3, 2, 3},
		{3, 6, 5},
		{4, 3, 3},
		{4, 5, 1},
		{5, 3, 1},
		{5, 6, 2},
	}

	dijkstra := NewFastDijkstra(nodeCount, edgeCount, graphRels)
	dijkstra.Execute(1)
	dijkstra.PrintResult()
}
