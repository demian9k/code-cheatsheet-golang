package algorithms

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	itr := []int{
		3, 6, 5, 7, 10, 8, 4, 2,
	}

	result := heapsort(itr, false)
	fmt.Println(result)
	//# [2, 3, 4, 5, 6, 7, 8, 10]

	result2 := heapsort(itr, true)
	fmt.Println(result2)
	//# [10, 8, 7, 6, 5, 4, 3, 2]
}
