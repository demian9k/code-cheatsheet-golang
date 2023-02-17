package gotest

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func Sum(a int, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	s := Sum(3, 6)

	if s != 9 {
		t.Error("Wrong result")
	}
}

/*
Benchmark** prefix를 가진다.
일반적으로 Benchmark prefix 다음 문자는 Uppercase로 시작한다.
*gotest.B 타입의 매개변수를 받습니다.
*/
func BenchmarkSum(b *testing.B) {
	/*
		goCommand$ go test -bench .
		.*나 .을 -bench argument 값으로 전달하면 모든 벤치마크 함수에 벤치마크 테스트를 수행

		루프의 수(b.N)는 Go의 테스트 도구가 적절한 값으로 지정
	*/
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

func divide(a, b int) {
	if b == 0 {
		panic("0으로 나눌 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

/*
assert 없이 panic 테스트
panic 테스트에선 defer recover가 panic보다 위에 있어야 한다.
*/
func TestPanicWithoutAssert(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	divide(1, 0)
}

/*
assert 를 사용한 panic 테스트
*/
func TestPanic(t *testing.T) {
	assert.Panics(t, func() { divide(1, 0) })
	assert.PanicsWithValue(t, "0으로 나눌 수 없습니다", func() { divide(1, 0) })
	assert.PanicsWithError(t, "error!", func() { panic(errors.New("error!")) })
}

/*
병렬 테스트
*/
func TestTeardownParallel(t *testing.T) {
	// <setup code>
	// This Run will not return until its parallel subtests complete.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			assert.True(t, false)
		})
		t.Run("Test2", func(t *testing.T) {
			assert.True(t, false)
		})
	})
	// <tear-down code>
}

/*
	Table driven test example
*/

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
		{input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for i, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
