package gotest

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
임의의 값 x를 매개변수로 받아 각 타입별 Len함수로 x의 길이를 반환하는 함수
from Go 언어 웹 프로그래밍 철저 입문 ( https://thebook.io/006806/ch07/03/06-03/ )
*/
func Len(x interface{}) int {
	value := reflect.ValueOf(x)
	switch reflect.TypeOf(x).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return value.Len()
	default:
		if method := value.MethodByName("Len"); method.IsValid() {
			values := method.Call(nil)
			return int(values[0].Int())
		}
	}

	panic(fmt.Sprintf("'%v' does not have a length", x))

}

func BenchmarkLenForString(b *testing.B) {
	b.StopTimer()
	v := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		//i를 string으로 변경
		v[i] = strconv.Itoa(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Len(v[i%1000000])
	}
}
