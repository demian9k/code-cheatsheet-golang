package algorithms

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

/*
계산 가능성 이론에 따르면 아커만 함수는 빌헬름 아커만(Wilhelm Ackermann)의 이름을 딴 함수로,
가장 간단하고 먼저 발견된 완전 계산 가능(total computable function) 함수
이며 비원시 귀납 함수(not primitive recursive function)이다.

모든 원시 귀납 함수는 완전(total)하고 계산 가능하지만, 아커만 함수는 모든 완전 계산 가능 함수가 원시 재귀 함수는 아니라는 것을 설명한다.
아커만의 발표 이후, 3개의 음수가 아닌 정수를 인자로 갖는 아커만 함수에 대해 많은 사람들이 다양한 목적으로 변형했다.

흔한 예정 하나로는 두 개의 인자를 갖는 아커만 피터 함수(Ackermann-Peter function)이 있는데 음수가 아닌 두 정수 m과 n에 대해 아래와 같이 정의한다.

아커만 함수의 극도로 깊은 재귀 성질 때문에 컴파일러의 재귀 최적화 성능 벤치마크를 측정하는 데 사용된다.

A(2,1) = A(1, A(2,0))
= A(1, A(1, A(2,-1)))
= A(1, A(1, A(1, A(2,-2))))
= A(1, A(1, A(1, A(1, A(2,-3)))))
= ...
= A(1, A(1, ... A(1, A(0, ... A(0,1))))))
= A(1, A(1, ... A(1, 2)))
= A(1, 3)
= A(0, A(1, 2))
= A(0, A(0, A(1, 1)))
= A(0, A(0, A(0, A(1, 0))))
= A(0, A(0, A(0, 2)))
= A(0, A(0, 3))
= A(0, 4)
= 5
*/
func Ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m > 0 && n == 0 {
		return Ackermann(m-1, 1)
	} else {
		return Ackermann(m-1, Ackermann(m, n-1))
	}
}

// 0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149
func Tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
	}
}

/*
루카스 수
Lucas numbers

피보나치 수와 같이 뒤의 두 항을 더하여 얻어진다.
2, 1, 3, 4, 7, 11, 18, 29, 47
*/
func Lucas(n int) int {
	if n == 0 {
		return 2
	} else if n == 1 {
		return 1
	} else {
		return Lucas(n-1) + Lucas(n-2)
	}
}

/*
조합론에서 빈번하게 나타나는 수
1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862
*/
func Catalan(n int) int {
	if n <= 1 {
		return 1
	} else {
		result := 0
		for i := 0; i < n; i++ {
			result += Catalan(i) * Catalan(n-i-1)
		}
		return result
	}
}

// 거듭제곱 함수 x^n
func Power(x, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else {
		return x * Power(x, n-1)
	}
}

/*
최대공약수
Greatest common divisor
*/
func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

//A, B, C 축 세개에
//A에 n개의 원반이 있다.

// hanoitower struct
type HanoiTower struct {
	count int
}

func (h *HanoiTower) hanoitower(n int, source string, aux string, target string) {
	h.count += 1

	if n == 1 {
		fmt.Println("Move disk 1 from", source, "to", target)
	} else {
		h.hanoitower(n-1, source, target, aux)
		fmt.Println("Move disk", n, "from", source, "to", target)
		h.hanoitower(n-1, aux, source, target)
	}
}

func (h *HanoiTower) Run(n int, source string, auxiliary string, target string) int {
	h.count = 0
	h.hanoitower(n, source, auxiliary, target)
	return h.count
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftHalf := MergeSort(arr[:mid])
	rightHalf := MergeSort(arr[mid:])

	i, j := 0, 0
	sorted := make([]int, 0, len(arr))

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] < rightHalf[j] {
			sorted = append(sorted, leftHalf[i])
			i++
		} else {
			sorted = append(sorted, rightHalf[j])
			j++
		}
	}

	sorted = append(sorted, leftHalf[i:]...)
	sorted = append(sorted, rightHalf[j:]...)
	fmt.Printf("\n MergeSort, %v", arr)
	return sorted
}

func MergeSortUsingCopy(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	leftHalf := MergeSortUsingCopy(arr[:mid])
	rightHalf := MergeSortUsingCopy(arr[mid:])

	return mergeUsingCopy(leftHalf, rightHalf)
}

func mergeUsingCopy(leftHalf, rightHalf []int) []int {
	result := make([]int, len(leftHalf)+len(rightHalf))
	i, j, k := 0, 0, 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] < rightHalf[j] {
			result[k] = leftHalf[i]
			i++
		} else {
			result[k] = rightHalf[j]
			j++
		}
		k++
	}
	//copy를 사용하는 방법은 append에 비해 비용이 더 든다.
	copy(result[k:], leftHalf[i:])
	copy(result[k+len(leftHalf)-i:], rightHalf[j:])

	return result
}
