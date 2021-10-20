package main

import "fmt"

// 문제
// fmt.Printf와 서식 문자를 이용한 코드 작성
// 이때 출력 결과의 최소 너비 6 지정해야함

func main() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Printf("%6d\n", a)   // 앞에 3칸 띄우고 123 출력
	fmt.Printf("%06d\n", b)  // 004567 출력
	fmt.Printf("%6.2f\n", f) // 앞에 두 칸 띄우고 3.14 출력
}
