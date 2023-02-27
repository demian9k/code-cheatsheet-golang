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

	var resultIndex = BinarySearch(nums, 22)

	if target != nums[resultIndex] {
		t.Errorf("target %d result %d not equal", target, nums[resultIndex])
	}
}
