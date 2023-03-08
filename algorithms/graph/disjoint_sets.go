package algorithms

import (
	. "code-cheatsheet-golang/datastructure"
	"fmt"
)

/*
서로소 집합

정의: 공통원소가 없는 두 집합

{1,2} 와 {3, 4} 두 집합은 공통원소가 없어서 서로소 관계이다.

서로소집합 자료구조는 서로소 부분 집합을 저장하고 있는 자료구조이다.

union과 find를 연산을 갖고 있다.

union 연산은 2개의 원소가 가진 집합을 하나로 합친다.
find 연산은 특정 원소가 속한 집합이 찾는다

서로소 집합 자료구조는 그래프를 위한 prerequisite 이다.

서로소 집합 알고리즘은
노드개수 V, find or union 연산 갯수 M일때 O(VM)

path compression 시
O(V + M(1 + log_2-M/V * V))

O(V + M(1 + math.log(V, 2-M/V) )

NUM: V
BASE: 2-M/V

log(a) / log(Base).

math.log(V, 2-M/V) = math.log(V) / math.log(2-M/V)
*/
type DisjointSets struct {
	v      int
	e      int
	parent []int
}

func (ds *DisjointSets) FindParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = ds.FindParent(parent, parent[x])
	}
	return parent[x]
}

func (ds *DisjointSets) Execute(nodePairs []NodePair) int {
	for i := 0; i < ds.e; i++ {
		a, b := nodePairs[i].FromNode, nodePairs[i].ToNode
		ds.UnionParent(ds.parent, a, b)
	}
	return 0
}

func (ds *DisjointSets) UnionParent(parent []int, a, b int) {
	a = ds.FindParent(parent, a)
	b = ds.FindParent(parent, b)

	//더 적은 숫자인 노드 번호가 부모가 된다.
	if a < b {
		ds.parent[b] = a
	} else {
		ds.parent[a] = b
	}
}

/*
		서로소 집합 사이클 판별
	    모든 간선을 하나씩 확인해 매 간선에 대해 union 또는 find 함수를 호출하여 동작한다.
	    Undirected Graph 에서만 적용 가능하다.
*/
func (d *DisjointSets) HasCycle(nodePairs []NodePair) bool {
	cycleExist := false

	for _, nodePair := range nodePairs {
		a, b := nodePair.FromNode, nodePair.ToNode // 하나의 edge는 2개의 노드가 연결되어 있다.

		// 그래프에 순환이 발생하는 조건, a노드의 부모와 b노드의 부모가 동일한 경우
		if d.FindParent(d.parent, a) == d.FindParent(d.parent, b) {
			cycleExist = true
			break
		} else {
			// 순환이 발생하지 않으면 union(합집합) 연산 수행
			d.UnionParent(d.parent, a, b)
		}
	}

	return cycleExist
}

func (d *DisjointSets) PrintResult() {
	resultSets := make([]int, 0, d.v)
	for i := 1; i <= d.v; i++ {
		resultSets = append(resultSets, d.FindParent(d.parent, i))
	}

	fmt.Println("각 원소가 속한 집합:")
	nodeIds := make([]int, d.v)
	for i := range nodeIds {
		nodeIds[i] = i + 1
	}
	matrix1 := [][]int{nodeIds, resultSets}
	fmt.Println(matrix1)

	fmt.Println()

	resultParents := make([]int, 0, d.v)
	fmt.Println("parent mapping TABLE")
	for i := 1; i <= d.v; i++ {
		resultParents = append(resultParents, d.parent[i])
	}

	fmt.Println(resultParents)
}

func NewDisjointSets(v, e int) *DisjointSets {
	//각 노드별 부모 테이블 초기화
	parent := make([]int, v+1)
	for i := 1; i <= v; i++ {
		parent[i] = i
	}
	return &DisjointSets{v, e, parent}
}
