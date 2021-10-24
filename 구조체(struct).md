## Struct (구조체)
- Go에서 struct는 Custom Data Type을 표현하는데 사용된다. struct는 필드들의 집합체이며 필드들의 컨테이너이다.
- struct는 필드 데이터만을 가지며, 메서드를 갖지 않는다. 메서드는 별도로 분리하여 정의한다. (Go Method)

- ### struct 선언
  - struct를 정의하기 위해서 type문을 사용한다. 아래 예시는 name과 age 필드를 갖는 person이라는 struct를 정의한 예시이다. 외부에서 사용하려면 Person으로 변경하면 된다.
  ```go
    package main
    
    import "fmt"
    
    type person struct {
      name string
      age int
    }
    func main() {
      p := person{}
      
      p.name = "Lee"
      p.age = 27
      fmt.Println(p) // { Lee 27 }
    }
  ```
- ### struct 객체 생성
  - 선언된 struct 타입으로부터 객체를 생성하는 몇 가지 방법들이 있다. 위 예시처럼 person{}을 이용해 빈 person 객체를 할당하고 나중에 채워넣을수도 있다.
  - struct 객체를 생성할때 초기값을 할당하는 방법도 있다. 필드명을 지정하지 않고 값을 넣을수도 있는데 그러면 만약 필드가 생략될 경우 생략된 필드들은 zero value를 갖는다.
  - new()를 이용해서도 객체를 생성할 수 있다. 모든 필드를 zero value로 초기화하고 객체의 포인터를 리던한다.
  - struct는 기본적으로 mutable 개체로서 필드값이 변화할 경우 해당 개체 메모리에서 직접 변경된다. 하지만 struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by value에 따라 객체를 복사해서 전달하게 된다. 그래서 만약 Pass by Reference로 struct를 전달하려면 struct의 포인터를 전달해야 한다.
  ```go
  var p1 person
  p1 = person{"Bob",20}
  p2 := person{name: "Sean", age:50}
  
  p := new(person)
  p.name = "Lee" // p가 포인터여도 . 을 사용한다.
  ```
  
- ### 생성자(constructor) 함수
  - 구조체의 필드가 사용전에 초기화 되어야 하는 경우도 있다. 만약 struct필드가 map타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct사용자가 매번 map을 초기화 해야 된다는 것을 기억할 필요가 없다. 이 경우에 생성자 함수를 사용한다. 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다.
  ```go
    package main
    
    type dict struct {
      data map[int]string
    }
    
    // 생성자 함수 정의
    func newDict() *dict {
      d := dict{}
      d.data = map[int]string{}
      return &d // 포인터 전달
    }
    
    func main() {
      dic := newDict() // 생성자 호출
      dic.data[1] = "A"
    }
  ```
