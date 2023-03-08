package shortest_path

import (
	"fmt"
)

type GraphRel struct {
	fromNode int
	toNode   int
	weight   int
}

type SimpleDijkstra struct {
	nodeCount int
	edgeCount int
	INF       int
	graph     [][]GraphRel
	visited   []bool
	distance  []int
}

func NewSimpleDijkstra(nodeCount, edgeCount int, graphRels []GraphRel) *SimpleDijkstra {
	s := &SimpleDijkstra{
		nodeCount: nodeCount,
		edgeCount: edgeCount,
		INF:       int(1e9),
		graph:     make([][]GraphRel, nodeCount+1),
		visited:   make([]bool, nodeCount+1),
		distance:  make([]int, nodeCount+1),
	}

	for _, rel := range graphRels {
		s.graph[rel.fromNode] = append(s.graph[rel.fromNode], rel)
	}

	return s
}

func (s *SimpleDijkstra) getSmallestNode() int {
	minVal := s.INF
	index := 0

	for i := 1; i <= s.nodeCount; i++ {
		if !s.visited[i] && s.distance[i] < minVal {
			minVal = s.distance[i]
			index = i
		}
	}

	return index
}

func (s *SimpleDijkstra) Execute(start int) {
	for i := 1; i <= s.nodeCount; i++ {
		s.distance[i] = s.INF
	}

	s.distance[start] = 0
	s.visited[start] = true

	for _, rel := range s.graph[start] {
		s.distance[rel.toNode] = rel.weight
	}

	for i := 1; i < s.nodeCount; i++ {
		now := s.getSmallestNode()
		s.visited[now] = true

		for _, rel := range s.graph[now] {
			cost := s.distance[now] + rel.weight

			if cost < s.distance[rel.toNode] {
				s.distance[rel.toNode] = cost
			}
		}
	}
}

func (s *SimpleDijkstra) PrintResult() {
	for i := 1; i <= s.nodeCount; i++ {
		if s.distance[i] == s.INF {
			fmt.Println("unreachable")
		} else {
			fmt.Println(s.distance[i])
		}
	}
}
