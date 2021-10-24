## 메서드(Method)
- go는 OOP를 고유의 방식으로 지원한다. 타 언어의 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go에서는 struct는 필드만을 가지고 메서드는 별도로 분리되어 정의된다.
- go 메서드는 특별한 형태의 func 함수이다. 메서드는 함수 정의에서 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다.
- 흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct변수명을 지정하는데, struct변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다.
```go
  package main
  // 1. Rect - struct 정의
  type Rect struct {
    width, height int
  }
  // 2. Rect의 area() 메서드
  func (r Rect) area() int {
    return r.width * r.height
  }
  
  func main() {
    rect := Rect{10,20}
    area := rect.area() // 3. 메서드 호출
    println(area) // 200
  }
```

- ### Value receiver vs 포인터 receiver
  - Value receiver는 struct의 데이터를 복사하여 전달하며, 포인터 receiver는 struct의 포인터만을 전달한다.
  - Value receiver의 경우 만약 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이터는 변경되지 않는다.
  - 포인터 receiver는 메서드 내의 필드값이 변경되면 그대로 호출자에 반영된다.
  - 위 예제는 Value receiver를 표현한것이라 area() 메서드 내에서 width나 height가 변경되더라도 main()함수의 rect구조체의 필드값에는 변화가 없다.
  - 아래 예제는 포인터 receiver로 변경후 메서드 내 r.width++필드 변경분이 main()함수에서도 반영되기 때문에 값이 11,220이 나온다.
  ```go
    package main
    
    type Rect struct {
      width, height int
    }
    
    // 포인터 receiver 사용 *Rect <-
    func (r *Rect) area2() int {
      r.width++
      return r.width * r.height
    }
    
    func main() {
      rect := Rect{10,20}
      area := rect.area2() // 메서드 호출
      println(rect.width, area) // 11 220 출력 
      // 만약 value receiver로 바꾸면 r.width++를 작성해도 main()함수의 rect구조체의 필드값에는 변화가 없음
    }
  ```
