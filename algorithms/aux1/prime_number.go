package algorithms

import "math"

// IsPrimeNumber1 가장 기본적인 소수 확인 알고리즘
func IsPrimeNumber1(x int) bool {
	// 2부터 (x -1) 까지의 모든 수를 확인하여
	for i := 2; i < x; i++ {
		// x가 해당 수로 나누어떨어진다면 소수가 아니다.
		if x%i == 0 {
			return false
		}
	}
	return true
}

/*
x의 약수가 대칭적인 특징을 이용해
O(n^0.5)로 줄인 알고리즘
*/
func IsPrimeNumber(x int) bool {
	// 2부터 x 의 제곱근까지 모든 수를 확인하여
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		// x가 해당 수로 나누어떨어진다면 소수가 아니다.
		if x%i == 0 {
			return false
		}
	}
	return true
}

/*
Sieve of Eratosthenes
에라토스테네스의 체
여러개의 수가 소수인지 아닌지를 판별할 때 사용하는 알고리즘
일정 범위의 수가 소수인지 확인할 때는 이 방법이 더 효율적인 알고리즘이다.

시간복잡도 O(NloglogN)

O(N) 만큼 빠르나 메모리가 많이 필요하다.
*/
func EratoSieve(n int) []int {
	booleanResult := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		booleanResult[i] = true
	}

	// 2부터 n 제곱근까지 모든 수를 확인
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		// i의 모든 배수 지우기
		j := 2
		for i*j <= n {
			booleanResult[i*j] = false
			j += 1
		}
	}

	primeNumbers := make([]int, 0)
	for i := 2; i <= n; i++ {
		if booleanResult[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
