# 배열(Array)
- 배열은 연속적인 메모리 공간에 동일한 타입의 데이터를 순차적으로 저장하는 자료구조이다. Go에서 배열의 첫번째 요소는 0,1,2,3.... 순으로 인덱스를 매긴다.
- 배열의 선언 형태는 "var 변수명[배열크기]데이터타입"과 같이 표현한다. 배열크기를 데이터타입 앞에 써주는것이 다른 언어들과는 다르다. 배열의 배열크기는 Type을 구성하는 한 요소이다. 예를 들어서 [3]int와 [5]int는 서로 다른 타입으로 인식된다. 배열이 선언된 후에 각 배열의 요소를 인덱스를 사용하여 읽거나 쓸 수 있다.
```go
package main

func main() {
  var a [3]int // 정수형 3개 요소를 갖는 배열 a를 선언한것
  a[0] = 1
  a[1] = 2
  a[2] = 3
  println(a[1]) // 2를 출력한다. ( a의 2번째 인덱스 )
```

- ### 배열의 초기화
  - 배열을 정의할 때, 초기값을 설정할 수 있다. 초기값은 "[배열크기]데이터타입"뒤에 {}괄호를 두고 초기값을 순서대로 작성하면 된다. 만약 초기화 과정에서 [...]를 사용해서 배열크기를 생략하면 자동으로 초기화 요소 숫자만큼 배열크기가 정해진다.
  ```go
    var a1 = [3]int{1,2,3}
    var a2 = [...]int{1,2,3} // 결국 위,아래 같은것임
  ```
- ### 다배열 배열
  - go는 다차원 배열을 지원한다. 다차원 배열은 배열크기 부분을 복수 개로 설정하여 선언한다. 예시로 3x4x5 차원 정수 배열을 만들면 아래와 같이 사용가능하다.
  ```go
    var multiArray [3][4][5] int
    multiArray[0][1][2] = 10
  ```
  
- ### 다차원 배열의 초기화
  - 다차원 배열의 초기화는 다차원 배열의 초기화와 비슷하다. 다만, 다차원이므로 배열 초기값 안에 다시 배열값을 넣는 형태를 띤다.
  ```go
    func main() {
      var a = [2][3]int{
        {1,2,3},
        {4,5,6},
      }
      println(a[1][2]) // 3
    }
  ```
  
  
# 슬라이스(Slice)
- go 배열은 고정된 배열크기 안에 동일한 타입의 데이터를 연속적으로 저장하지만, 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능을 가지고 있지 않다.
- go slice는 내부적으로 배열에 기초하여 만들어 졌지만 배열의 이런 제약점들을 넘어 개발자에게 편리하고 유용한 기능들을 제공한다.
- 슬라이스는 배열과 달리 고정된 크기를 미리 지정하지 않을수 있고, 차후 그 크기를 동적으로 변경할 수도 있고, 부분 배열을 발췌할 수도 있다.
- go slice 선언은 배열을 선언하듯 "var v []T"형태로 하는데 배열과 달리 크기는 지정하지 않는다.
```go
package main
import "fmt"

func main() {
  var a[] int // 슬라이스 변수 선언
  a = []int{1,2,3} // 슬라이스에 리터럴값 지정
  a[1] = 10
  fmt.Println(a) // [1,10,3] 나옴
```

- ### make()
  - go에서 slice를 생성하는 다른 방법으로 내장함수인 make()를 사용할 수 있다.
  - make()함수로 슬라이스를 생성하면, 개발자가 슬라이스의 길이와 용량을 임의로 지정할 수 있다. 
  ```go
    func main() {
      arr := make([]int, 5, 10) // 1. 타입지정, 2. 길이, 3. 내부 배열의 최대 길이
      println(len(s), cap(s)) // len 5, cap 10
  ```
  ```go
    // 이렇게 별도의 길이,용량 없이 슬라이스를 만들면 기본적으로 길이,용량 0으로 생성된다.
    // 이를 Nil Slice라고 하고 nil과 비교하면 참을 리턴함
    func main() {
      var s []int
      
      if s == nil {
        println("Nil Slice")
      }
      println(len(s), cap(s)) // 0,0
  ```
- ### 부분 슬라이스(sub-slice)
  - 슬라이스에서 일부를 발췌하여 부분 슬라이스를 만들수있다. "슬라이스[처음인덱스:마지막인덱스]"형식으로 만든다.
  - 처음,마지막을 생략 가능하다.
  ```go
    package main
    import "fmt"
    
    func main() {
      s := []int{0,1,2,3,4,5}
      s = s[2:5] // 2번째 인덱스부터 5번째 인덱스의 앞까지 나온다. 고로 2,3,4번째 인덱스들이 나옴
      fmt.Println(s) // 2,3,4가 출력된다. 
  ```
  ```go
    s := []int{0,1,2,3,4,5}
    s = s[2:5] // 2,3,4
    s = s[1:] // 3,4 
    fmt.Println(s) // 3,4
  ```
  
- ### 슬라이스 병합
  - 배열은 고정된 크기로 크기 이상의 데이터를 임의로 추가할 수 없다. 하지만 슬라이스는 자유롭게 새로운 요소를 추가할 수 있다.
  - go의 내장함수인 append()를 사용한다. 첫 번째 파라미터는 슬라이스 객체이고, 두 번째는 추가할 값이다.
  ```go
    func main() {
      s := []int{0,1}
      
      s = append(s, 2) // 1개 추가해서 이제 0,1,2
      s = appned(s, 3,4,5) // 3개 추가해서 이제 0,1,2,3,4,5
      fmt.Println(s) // 0,1,2,3,4,5
  ```
  
# Map
- Map은 키(key)에 대응하는 값(value)를 신속히 찾는 해시테이블을 구현한 자료구조이다. go는 Map타입을 내장하고 있는데, "map[key타입]value타입"과 같이 선언할 수 있다.
```go
// 정수를 key로 하고 문자열을 값으로 하는 맵 변수이다.
var idMap map[int]string
// idMap은 nil값을 갖고 이를 Nil Map이라 부른다. NilMap은 어떤 데이터를 쓸 수 없는데, map을 초기화하기 위해 make()함수를 사용할 수 있다.
idMap = make(map[int]string)
// 리터럴을 사용해서 초기화할 수도 있다.
// 리터럴 초기화는 "map[key타입]value타입 {key:value}"와 같이 Map타입 뒤 {}괄호 안에 "키:값"들을 나열하면 된다.
tickers := map[int]string{
  "GOOD": "Google Inc",
  "MSFT": "Microsoft",
  "FB": "Facebook",
}
```
- ### Map 사용
  - 처음 map이 초기화 되었을때 아무 데이터가 없는 상태이다. 이때 새로운 데이터를 추가하기 위해서는 "map변수[키] = 값"과 같이 해당 키에 그 값을 할당하면 된다.
  - 만약 key 901에 Apple을 할당하면 새 해시 키-값 쌍이 추가된다. 만약 901이 이미 있다면 값만 갱신한다.
  ```go
    package main
    
    func main() {
      var m map[int]string
      
      m = make(map[int]string)
      // 추가 및 갱신
      m[901] = "Apple"
      m[134] = "Grape"
      m[777] = "Tomato"
      
      // 키에 대한 값 읽기
      str := m[134]
      println(str) // Grape
      
      // 값이 없으면 nil 또는 zero 리턴한다
      noData := m[999]
      println(noData) // nil or zero
      
      // 삭제하기
      delete(m, 777)
  ```
- ### Map 키 체크
  - map안에 특정 키가 존재하는지 체크할 필요가 가끔 있는데, 이를 위해 Go에선 "map변수[키]"읽기를 수행할 때 2개의 리턴값을 리턴한다. 
  - 첫 번째는 키에 상응하는 값이고, 두 번째는 그 키가 존재하는지 아닌지를 나타내는 bool값이다.
  ```go
    package main
    
    func main() {
      tickers := map[string]string{
        "GOOD": "Google Inc",
        "MSFT": "Microsoft",
        "FB": "Facebook",
        "AMZN": "Amazon",
      }
      
      // map 키 체크하기
      val, exists := tickers["MSFT"]
      if !exists { // "MSFT"키는 존재하기 때문에 exists는 false라서 이건 출력안된다.
        println("No MSFT ticker")
      }
      println(val, exists) // Microsoft true
    }
  ```
  
- ### for 루프를 사용한 Map 열거
  - Map이 갖고 있는 모든 요소들을 출력하기 위해, for range 루프를 사용할 수 있다. Map 컬렉션에 for range를 사용하면, Map키와 Map 값 2개의 데이터를 리턴한다.
  ```go
    package main
    
    import "fmt"
    
    func main() {
      myMap := map[string]string {
        "A": "Apple",
        "B": "Banana",
        "C": "Charlie",
      }
      
      for key, val := range myMap {
        fmt.Println(key, val)
      }
    }
  ```
