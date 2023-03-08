package algorithms

import (
	. "code-cheatsheet-golang/datastructure"
	"container/list"
)

/*
위상 정렬

정의: Directed Graph의 모든 노드를 방향성에 거스르지 않도록 순서대로 나열하는 것

시간 복잡도
V = vertex = 노드 갯수
E = edge = 간선 갯수
O(V+E)
*/
func TopologySort(v int, nodePairs []*NodePair) []int {
	// 진입 차수 0으로 초기화
	indegree := make([]int, v+1)
	graph := make([][]int, v+1)

	// 방향 그래프의 모든 간선 정보를 입력받기
	for _, nodePair := range nodePairs {
		graph[nodePair.FromNode] = append(graph[nodePair.FromNode], nodePair.ToNode)

		indegree[nodePair.ToNode]++ // toNode id로 테이블에 진입차수를 1로 증가시킨다.
	}

	//정렬된 node id 목록
	var result []int

	q := list.New()

	//처음 시작할 때는 진입차수가 0인 노드를 큐에 삽입
	for i := 1; i <= v; i++ {
		if indegree[i] == 0 {
			q.PushBack(i)
		}
	}

	//큐가 빌 때까지 반복
	for q.Len() > 0 {
		now := q.Front().Value.(int)
		q.Remove(q.Front())
		result = append(result, now)

		// 해당 원소와 연결된 노드들의 진입차수에서 1 빼기
		for _, i := range graph[now] {
			indegree[i]--

			//새롭게 진입차수가 0이 되는 노드가 있다면 큐에 삽입해서 while문을 계속 진행
			if indegree[i] == 0 {
				q.PushBack(i)
			}
		}
	}

	return result
}
