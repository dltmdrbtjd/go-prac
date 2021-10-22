package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	// n개의 개수 스캔

	slice := make([]int, 0, n)
	// go에선 slice만드는거라고 하는데
	// 배열과 같긴한데 길이가 고정되어 있지 않고 동적으로 크기가 늘어남
	// 위 예제는 int형 slice를 선언한거고 길이는 0이고 용량이 n인거임

	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		// 이제 n개의 수들을 스캔해서 slice 넣어줘야함
		slice = append(slice, num)
		// append를 사용하면 흠.. 기존 slice에 새로운 num을 추가하는 방식으로 보여지는데
		// 정확한지는 모르겠다.
	}

	sort.Ints(slice)
	// https://pkg.go.dev/sort 여기 공식문서에 설명되어있음
	// 이제 정렬을 위해서 sort패키지 사용해야하는데
	// 그중 Ints라는 함수 사용하면 배열들을 오름차순으로 정렬시켜줌.

	for _, i := range slice {
		fmt.Println(i)
		// 이제 정렬된 slice를 반복문을 사용해서 값을 찍어주면된다.
		// _는 키 값이니까 언더스코어로 사용안하게 처리하고
		// i로 값을 받아서 프린트 쭉쭉 찍어주면 끝
	}
}
