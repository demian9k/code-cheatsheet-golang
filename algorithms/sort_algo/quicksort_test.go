package sort_algo

import (
	"testing"
)

func TestQuickSort(t *testing.T) {

	nums := []int{
		16, 100, 44, 94, 22, 23, 77, 11, 22, 33,
		44, 55, 66, 99, 12, 1, 2, 3, 4, 5,
		108,
	}
	var sortedNums = quickSort(nums, 0, len(nums)-1)
	//var sortedNums = nums
	for i, v := range sortedNums {
		if i > 0 && sortedNums[i-1] > v {
			//fmt.Println(i, v)
			t.Errorf("%d %d error", sortedNums[i-1], v)
			break
		}
	}
}
