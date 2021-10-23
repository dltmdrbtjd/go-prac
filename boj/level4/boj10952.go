package main

import "fmt"

func main() {
	var a, b int
	// int값 a,b선언

	// go에는 while문법이 없음 그래서 for true로 사용가능
	// 그냥 for 만 써도 괜찮음
	for true {
		// 입력받은 값 스캔
		fmt.Scan(&a, &b)
		// 마지막 값이 항상 0 0이니까 그럴경우에 break걸어주기
		if a == 0 && b == 0 {
			break
		} else {
			// 0 0 이 아니면 더한값 출력  0 0이 될때까지 !
			fmt.Println(a + b)
		}
	}
}
