package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanln(a, b)
	// ScanIn()인수로 변수의 메모리 주소를 넘겨야함.
	// 그런데 그냥 변수를 넘겨버려서 제대로 동작 안함
	// fmt.ScanIn($a,$b)로 수정해야함.
	fmt.Println(a, b)
}
