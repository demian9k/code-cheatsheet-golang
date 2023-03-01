package algorithms

import (
	"fmt"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	nums := []int{
		16, 100, 44, 94, 22, 23, 77, 11, 22, 33,
		55, 66, 99, 12, 1, 2, 3, 4, 5,
		108,
	}
	fmt.Println(nums)

	//sort.Slice는 reflect 사용
	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})

	sort.Ints(nums)

	target := nums[17] //99

	var resultIndex = BinarySearch(nums, target, 0, len(nums)-1)

	if target != nums[resultIndex] {
		t.Errorf("target %d result %d not equal", target, nums[resultIndex])
	}
}

func TestBinarySearch_recur(t *testing.T) {

	nums := []int{
		1, 2, 3, 4, 5, 11, 12, 16, 22, 23, 24, 33, 44, 55, 66, 77, 94, 99, 100, 108,
	}
	fmt.Println(nums)

	test1 := func(target int) {
		//target := nums[targetIndex] //99
		resultIndex := BinarySearch_recur(nums, target, 0, len(nums)-1)

		if target != nums[resultIndex] {
			t.Errorf("target %d result %d not equal", target, nums[resultIndex])
		}
	}

	test1(nums[17])
	test1(nums[7])
}

func TestBinarySearchTailRecursion(t *testing.T) {

	nums := []int{
		1, 2, 3, 4, 5, 11, 12, 16, 22, 23, 24, 33, 44, 55, 66, 77, 94, 99, 100, 108,
	}
	fmt.Println(nums)

	test1 := func(target int) {
		//target := nums[targetIndex] //99
		resultIndex := BinarySearchTailRecursion(nums, target)

		fmt.Println(resultIndex)

		if target != nums[resultIndex] {
			t.Errorf("target %d result %d not equal", target, nums[resultIndex])
		}
	}

	test1(nums[17])
	test1(nums[7])
}

func BenchmarkBinarySearchRecur(b *testing.B) {
	nums := []int{
		1, 2, 3, 4, 5, 11, 12, 16, 22, 23, 24, 33, 44, 55, 66, 77, 94, 99, 100, 108,
	}

	for i := 0; i < b.N; i++ {
		BinarySearch_recur(nums, nums[17], 0, len(nums)-1)
	}
}

func BenchmarkBinarySearchTailRecur(b *testing.B) {
	nums := []int{
		1, 2, 3, 4, 5, 11, 12, 16, 22, 23, 24, 33, 44, 55, 66, 77, 94, 99, 100, 108,
	}

	for i := 0; i < b.N; i++ {
		BinarySearchTailRecursion(nums, nums[17])
	}

	//for i := 0; i < b.N; i++ {
	//	BinarySearch_recur(nums, nums[17], 0, len(nums)-1)
	//}
}
