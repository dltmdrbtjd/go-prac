package main

import "fmt"

func Add(a int b int) int {
	return a + b
}

// 1. 함수 호출
// 2. 매개변수 생성 및 초기화
// 3. a/b더한 값 반환(복사)
// 4. c에 복사
// 5. 로컬 변수 삭제

func main() {
	c := Add(6,7)
	fmt.Println(c)
}