package main

func sum(a []int) int {
	var r int
	for _, i := range a {
		r += i
	}
	return r
}

// https://knight76.tistory.com/entry/go-lang-underscore-%EC%9D%98-%EC%82%AC%EC%9A%A9
// underscore가 key를 사용하지 않는다는 의미가 있고
// 값을 대입받는 의미로도 사용된다.
// 그럼 _는 키 인거고 i가 값인걸까??
// 그럼 만약 [1,2,3,4,5] 인거면
// 0:1
// 1:2
// 2:3
// 3:4
// 4:5
// 이렇게 나와서 _는 키 인거고
// i에는 값들이 들어가는거겠지 ??
// 아 그러면 저 문법이 이해가되네
