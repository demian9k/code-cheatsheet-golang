package shortest_path

import (
	"code-cheatsheet-golang/datastructure/heapq"
	"container/heap"
	"fmt"
)

// 구현이 어렵고 빠른 코드 (E log V)
type FastDijkstra struct {
	nodeCount int
	edgeCount int
	INF       int
	graph     [][]GraphRel
	distance  []int
}

func NewFastDijkstra(nodeCount, edgeCount int, graphRels []GraphRel) *FastDijkstra {
	dijkstra := &FastDijkstra{
		nodeCount: nodeCount,
		edgeCount: edgeCount,
		INF:       1e9,
		graph:     make([][]GraphRel, nodeCount+1),
		distance:  make([]int, nodeCount+1),
	}

	for i := 0; i <= nodeCount; i += 1 {
		dijkstra.distance[i] = 1e9
	}

	for _, rel := range graphRels {
		dijkstra.graph[rel.fromNode] = append(dijkstra.graph[rel.fromNode], rel)
	}

	return dijkstra
}

func (d *FastDijkstra) Execute(start int) {
	q := make(heapq.PriorityQueue[int], 0)

	//시작 노드로 가기 위한 최단 경로는 0으로 설정하여, 큐에 삽입
	//heapq를 일반 리스트를 파라미터로 넣으면 최소힙 우선순위 큐로 취급된다.
	heap.Push(&q, &heapq.Item[int]{Priority: 0, Value: start})
	d.distance[start] = 0

	for len(q) > 0 {

		// 가장 최단 거리가 짧은 노드에 대한 정보 꺼내기
		// get_smallest_node() 메서드를 자료구조로 대체하는 것
		item := heap.Pop(&q).(*heapq.Item[int])
		dist, now := item.Priority, item.Value

		// 거리 테이블에 거리가 최소인 node_id로 접근하여 이전에 계산된 거리값과 현재 dist 비교
		// 가장 작은 현재 노드가 이미 처리된 적이 있는 노드라면 무시
		if d.distance[now] < dist {
			continue
		}

		// 현재 노드와 연결된 다른 인접한 노드들을 확인
		for _, rel := range d.graph[now] {
			cost := dist + rel.weight

			//현재 노드를 거쳐서, 다른 노드로 이동하는 거리가 더 짧은 경우
			if cost < d.distance[rel.toNode] {
				d.distance[rel.toNode] = cost
				//heapq에 튜플을 넣으면 첫번째 값을 기준으로 정렬한다.
				heap.Push(&q, &heapq.Item[int]{Value: rel.toNode, Priority: cost})
			}
		}

	}
}

func (d *FastDijkstra) PrintResult() {
	for i := 1; i <= d.nodeCount; i++ {
		if d.distance[i] == d.INF {
			fmt.Println("unreachable")
		} else {
			fmt.Println(d.distance[i])
		}
	}
}
