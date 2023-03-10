package datastructure

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFenwickTree_Build(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	arr := produceOriginalArray()
	n := len(arr)
	ft := NewFenwickTree(n)

	// 트리 만들기
	for i := 0; i < n; i++ {
		ft.Update(i, arr[i])
	}

	t.Log(ft.tree)
}

func TestFenwickTree_UpdateAndGet(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	arr := produceOriginalArray()
	n := len(arr)
	ft := NewFenwickTree(n)

	// 트리 만들기
	for i := 0; i < n; i++ {
		ft.Update(i, arr[i])
	}

	t.Log(ft.tree)

	// 0~7 까지 prefix sum
	sum := ft.Query(7)
	t.Log(sum)
}

func produceOriginalArray() []int {
	return []int{3, 5, 2, 4, 0, 8, 1, 5, 3, 4, 1, 0, 6, 7, 7, 3}
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Printf("\033[1;34m%s\033[0m", ">> Setup Test\n")

	return func(tb testing.TB) {
		log.Printf("\033[1;34m%s\033[0m", ">> Teardown Test\n")
	}
}

func setupMain() {
	// Do something here.
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")
}

func teardownMain() {
	// Do something here.
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setupMain()
	code := m.Run()
	teardownMain()
	os.Exit(code)
}
