package algorithms

import (
	. "code-cheatsheet-golang/datastructure"
	"sort"
)

/*
NewKruskalAlgorithm

신장 트리 ( spanning tree )

정의: 하나의 그래프에서 모든 노드를 포함하지만 순환이 없는 부분 그래프

크루스칼(Kruskal) 알고리즘

하나의 그래프가 있을 때 최소 비용을 갖는 신장 트리를 찾는 알고리즘

시간 복잡도
O(ElogE)
*/
func NewKruskalAlgorithm(v int, e int) *KruskalAlgorithm {
	return &KruskalAlgorithm{
		ds: NewDisjointSets(v, e),
	}
}

type KruskalAlgorithm struct {
	ds *DisjointSets
}

func (k *KruskalAlgorithm) Execute(graphRels []*GraphRel) int {
	sort.Slice(graphRels, func(i, j int) bool {
		return graphRels[i].Weight < graphRels[j].Weight
	})

	totalCosts := 0

	// edge를 하나씩 루프하며 사이클을 발생시키는지 확인한다.
	// 발생하면 포함시키지 않는다.
	for _, edge := range graphRels {
		if k.ds.FindParent(k.ds.parent, edge.FromNode) != k.ds.FindParent(k.ds.parent, edge.ToNode) {
			// 사이클이 발생하지 않으면 트리에 포함시킨다.
			k.ds.UnionParent(k.ds.parent, edge.FromNode, edge.ToNode)
			totalCosts += edge.Weight // 비용은 합산해준다.
		}
	}

	return totalCosts
}
