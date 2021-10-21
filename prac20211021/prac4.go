package main

import "fmt"

func F(n int) int {
	// 탈출 조건
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	// 피보나치 수열 9 번째 값
	fmt.Println(F(9))
}
