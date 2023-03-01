package perf

import (
	"testing"
)

// tail-recusion 과 memoization을 사용한 fibonacci 재귀 함수
func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibo(20)
	}
}

// tail-recusion을 사용한 fibonacci 재귀 함수
func BenchmarkTailFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboTr(20)
	}
}

// memoization을 사용한 fibonacci 재귀 함수
func BenchmarkMemoFibo(b *testing.B) {
	//fibo := ProduceFiboMemo()
	for i := 0; i < b.N; i++ {
		//fibo(20)
		ProduceFiboMemo()(20)
	}
}

// tail-recusion 과 memoization을 사용한 fibonacci 재귀 함수
func BenchmarkTailRecMemoFibo(b *testing.B) {
	//fibo_tr_memo_helper := produceFiboTrMemo()
	for i := 0; i < b.N; i++ {
		//fibo_tr_memo_helper(20, 1, 1)
		FiboTrMemo(20)
	}
}
