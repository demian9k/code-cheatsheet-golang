package algorithms

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	/*
	 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144.
	*/

	answer := Fibonacci(6)
	if answer != 8 {
		t.Errorf("Expected 8, but got %d", answer)
	}
}

func TestFactorial(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	/*
	   5 x 4 x 3 x 2 x 1
	*/
	answer := Factorial(5)
	if answer != 120 {
		t.Errorf("Expected 120, but got %d", answer)
	}
}

func TestAckermann(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	/*
	   2,1
	*/
	answer := Ackermann(2, 1)
	if answer != 5 {
		t.Errorf("Expected 5, but got %d", answer)
	}
}

func TestTribonacci(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	answers := []int{0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890}

	n := 15
	answer := Tribonacci(n)
	if answer != answers[n] {
		t.Errorf("Expected %d, but got %d", answers[n], answer)
	}
}

func TestLucas(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	// 2, 1, 3, 4, 7, 11, 18, 29, 47
	answer := Lucas(6)
	if answer != 18 {
		t.Errorf("Expected 18, but got %d", answer)
	}
}

func TestCatalan(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	answers := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796, 58786}

	n := 10
	answer := Catalan(n)
	if answer != answers[n] {
		t.Errorf("Expected %d, but got %d", answers[n], answer)
	}
}

func TestPower(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	// 지수함수
	answer := Power(2, 10)
	if answer != 1024 {
		t.Errorf("Expected 1024, but got %d", answer)
	}
}

func TestGCD(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	answer := GCD(24, 36)
	if answer != 12 {
		t.Errorf("Expected 12, but got %d", answer)
	}
}

func TestHanoi(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	n := 10
	tower := &HanoiTower{}
	result := tower.Run(n, "A", "C", "B")
	answer := int(math.Pow(2, float64(n))) - 1
	if result != answer {
		t.Errorf("Expected %d, but got %d", answer, result)
	}
}

func TestMergeSort(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	answer := MergeSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(answer, expected) {
		t.Errorf("MergeSort was incorrect, got: %v, want: %v", answer, expected)
	}
}

func TestMergeSortUsingCopy(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	answer := MergeSortUsingCopy([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	//expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\n TestMergeSort2, %v \n", answer)

	//if !reflect.DeepEqual(answer, expected) {
	//	t.Errorf("MergeSort was incorrect, got: %v, want: %v", answer, expected)
	//}
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
