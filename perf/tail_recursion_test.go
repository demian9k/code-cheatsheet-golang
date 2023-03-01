package perf

import "testing"

// tail-recusion 없는 팩토리얼 재귀 함수 테스트
// 각 호출마다의 stackframe을 관리해야 하므로 tail-rec 보다 부하가 있고
// 가장 깊은 함수의 호출이 모두 끝날때까지 함수 호출이 끝나지 않는 구조다
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(20)
	}
}

// tail-recusion 을 사용한 팩토리얼 재귀 함수 테스트
// stackframe들의 호출 부담이 없으며 매번 함수호출이 완료되고 다음 호출로 이어지는 구조이다.
func BenchmarkFactorialTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tailFact(20)
	}
}
