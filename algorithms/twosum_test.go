package algorithms

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	arr := []int{
		16, 100, 44, 94, 22, 23, 77, 11, 22, 33,
		44, 55, 66, 99, 12, 1, 2, 3, 4, 5,
		108,
	}

	target := 208 //208
	//arr에서 2수를 합쳐 target 이 되는 수 찾기
	var result = TwoSum(arr, target)

	sum := 0

	if len(result) == 2 {
		sum = arr[result[0]] + arr[result[1]]
	} else {
		t.Errorf("result is not pair")
	}

	if target != sum {
		t.Errorf("sum is not equal to target")
	}
}
