package main

import "fmt"

func main() {
	var a int8 = 30

	a <<= 2
	// a <<= 2는 a = << 2 와 같음
	// 즉 a의 비트를 왼쪽 2번 이동시킴 4배한 값과 같아짐 120이 된다.
	a += 8
	// a += 8 은 a = a+8과 같음
	// 고로 a = 128인데 int8 최대값이 127이라 -128로 출력
	fmt.Println(a)
}
