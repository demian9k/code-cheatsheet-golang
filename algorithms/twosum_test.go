package algorithms

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{
		16, 100, 44, 94, 22, 23, 77, 11, 22, 33,
		44, 55, 66, 99, 12, 1, 2, 3, 4, 5,
		108,
	}
	target := 185
	var indices = TwoSum(nums, target)

	sum := 0

	for _, i := range indices {
		sum += nums[i]
	}

	if target != sum {
		t.Errorf("sum is not equal to target")
	}
}
