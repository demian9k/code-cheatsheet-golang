package algorithms

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

// 구현이 쉽고 느린 코드
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

	// node id(번호)를 모두 loop한다.
	// distance list와 visited list를 접근하는 키로 사용된다.
	for i := 1; i <= s.nodeCount; i++ {
		// 방문하지 않은 노드중 현재 최솟값보다 작으면 최솟값 갱신
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
	// 시작노드 초기화
	s.distance[start] = 0
	s.visited[start] = true

	//distance 테이블 초기화
	for _, rel := range s.graph[start] {
		s.distance[rel.toNode] = rel.weight
	}
	// 시작 노드를 제외한 전체 n - 1개의 노드에 대해 반복
	for i := 1; i < s.nodeCount; i++ {
		// 거리가 가장 짧은 노드를 꺼내서 방문한 것으로 표시
		now := s.getSmallestNode()
		s.visited[now] = true

		for _, rel := range s.graph[now] {
			cost := s.distance[now] + rel.weight
			// 현재 노드를 거쳐서 다른 노드로 이동하는 거리가 더 짧은 경우
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
