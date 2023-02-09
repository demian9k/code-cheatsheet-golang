package gotest

import "fmt"

/*
Example로 시작되는 경우도 test 함수이다.
*/
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

/*
Example에서 Output: 이후 문자열과 println의 값이 같으면 test를 PASS 한다.
*/
func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
