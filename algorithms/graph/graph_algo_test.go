package algorithms

import (
	. "code-cheatsheet-golang/datastructure"
	"code-cheatsheet-golang/util"
	"fmt"
	"testing"
)

func TestDisjointSets_Execute(t *testing.T) {
	disjointSets := NewDisjointSets(6, 4)
	nodePairs := []NodePair{{1, 4}, {2, 3}, {2, 4}, {5, 6}}
	disjointSets.Execute(nodePairs)
	disjointSets.PrintResult()
}

func TestDisjointSets_HasCycle(t *testing.T) {
	nodePairs := []NodePair{{1, 2}, {1, 3}, {2, 3}}
	disjointSets := NewDisjointSets(3, 3)
	hasCycle := disjointSets.HasCycle(nodePairs)

	if hasCycle {
		fmt.Println("graph has cycle")
	} else {
		fmt.Println("graph has not cycle")
	}
}

func TestKruskalAlgorithm(t *testing.T) {
	edges := []*GraphRel{
		NewGraphRel(1, 2, 29),
		NewGraphRel(1, 5, 75),
		NewGraphRel(2, 3, 35),
		NewGraphRel(2, 6, 34),
		NewGraphRel(3, 4, 7),
		NewGraphRel(4, 6, 23),
		NewGraphRel(4, 7, 13),
		NewGraphRel(5, 6, 53),
		NewGraphRel(6, 7, 25),
	}

	nodes := []int{}
	for _, e := range edges {
		nodes = append(nodes, e.ToNode, e.FromNode)
	}

	uniqueNodes := util.RemoveDuplicates(nodes)

	kru := NewKruskalAlgorithm(len(uniqueNodes), len(edges))
	totalCosts := kru.Execute(edges)
	fmt.Println(totalCosts)
}

func TestTopologySort(t *testing.T) {
	nodePairs := []*NodePair{
		NewNodePair(1, 2),
		NewNodePair(1, 5),
		NewNodePair(2, 3),
		NewNodePair(2, 6),
		NewNodePair(3, 4),
		NewNodePair(4, 7),
		NewNodePair(5, 6),
		NewNodePair(6, 4),
	}

	nodes := []int{}
	for _, e := range nodePairs {
		nodes = append(nodes, e.ToNode, e.FromNode)
	}

	uniqueNodes := util.RemoveDuplicates(nodes)

	result := TopologySort(len(uniqueNodes), nodePairs)

	//큐에 새로 들어가는 원소가 2개 이상인 경우 답은 여러개가 될 수 있다.
	// 1,2,5,3,6,4,7
	// 1,5,2,3,6,4,7
	for _, i := range result {
		fmt.Printf("%d ", i)
	}
}
