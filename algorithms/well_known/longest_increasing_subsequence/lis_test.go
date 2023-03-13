package algorithms

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Printf("\033[1;34m%s\033[0m", ">> Setup Test\n")

	return func(tb testing.TB) {
		log.Printf("\033[1;34m%s\033[0m", ">> Teardown Test\n")
	}
}

func produceOriginalArray() []int {
	return []int{10, 20, 10, 30, 20, 50}
}

func TestLISDPMaxLen(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	arr := produceOriginalArray()
	maxLen := GetLongestIncreasingSubsequence(arr)
	answer := 4
	if maxLen != 4 {
		t.Errorf("%d != %d error", answer, maxLen)
	}
}

func TestLISStruct(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	arr := produceOriginalArray()
	lisInstace := NewLIS(arr)

	maxLen := lisInstace.GetLISLength()

	answer := 4
	if maxLen != 4 {
		t.Errorf("%d != %d error", answer, maxLen)
	}

	lis := lisInstace.GetLIS()

	fmt.Println(lis)
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
