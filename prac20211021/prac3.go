package main

import "fmt"

// 함수 이름은 Mutiple이다.
// 입력으로 int 타입 2개를 받고 int 타입값 1개를 반환한다.
// 두 입력값을 곱한 결과값을 반환한다.

func Multiple(a int b int) int {
	return a * b
}

func main() {
	seven := Multiple(6,6)
	println(seven)
}
