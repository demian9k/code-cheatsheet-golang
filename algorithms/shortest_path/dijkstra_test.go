package algorithms

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNewSimpleDijkstra(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	graphRels, nodeCount, edgeCount := produceGraphRelsData()

	dijkstra := NewSimpleDijkstra(nodeCount, edgeCount, graphRels)
	dijkstra.Execute(1)
	dijkstra.PrintResult()
}

func TestNewFastDijkstra(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	graphRels, nodeCount, edgeCount := produceGraphRelsData()

	dijkstra := NewFastDijkstra(nodeCount, edgeCount, graphRels)
	dijkstra.Execute(1)
	dijkstra.PrintResult()
}

func produceGraphRelsData() ([]GraphRel, int, int) {
	return []GraphRel{
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
	}, 6, 11
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Printf("\033[1;34m%s\033[0m", ">> Setup Test\n")

	return func(tb testing.TB) {
		log.Printf("\033[1;34m%s\033[0m", ">> Teardown Test\n")
	}
}

func setupMain() {
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")
}

func teardownMain() {
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setupMain()
	code := m.Run()
	teardownMain()
	os.Exit(code)
}
