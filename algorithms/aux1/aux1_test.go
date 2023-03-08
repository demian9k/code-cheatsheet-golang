package algorithms

import (
	"reflect"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	var result = IsPrimeNumber1(7)

	if result == true {
	} else {
		t.Errorf("error")
	}

	var result2 = IsPrimeNumber(7)

	if result2 == true {
	} else {
		t.Errorf("error")
	}
}

func TestEratoSieve(t *testing.T) {
	var result = EratoSieve(120)
	answer := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
	//fmt.Println(result2)
	if reflect.DeepEqual(result, answer) {

	} else {
		t.Errorf("error")
	}
}
