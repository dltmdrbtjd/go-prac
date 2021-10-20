package main

import "fmt"

func main() {
	var a int32 = 360
	var b int8 = int8(a)

	//  b는 int8 type으로 -128~127 까지 값이 표현가능
	// 360은 int8의 범위를 벗어나서 마지막 1바이트 값만 남고 나머지가 사라져서 올바르지 않은 값 출력
	// 마지막 1byte값이 104라는건데... 1101000 ?

	fmt.Println(b)
}
