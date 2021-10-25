## 간단한 HTTP 서버
- go의 표준 패키지인 net/http 패키지는 웹 관련 서버(및 클라이언트) 기능을 제공한다.
- HTTP 서버를 만들기 위해 중요한 http 패키지 메서드로 ListenAndServe(), Handle(), HadleFunc()등이 있다.
- ListenAndServe() 메서드는 지정된 포트에 웹 서버를 열고 클라이언트 Request를 받아들여 새 Go 루틴에 작업을 할당하는 일을 한다.
- Handle()과 HandleFunc()메서드는 요청된 Request Path에 어떤 Request 핸들러를 사용할 지를 지정하는 라우팅 역할을 한다.
```go
package main

import "net/http"

func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello World"))
  })
  http.ListenAndServe(":5000", nil)
}
```
1. http://localhost:5000/hello로 접근한다.
2. HandleFuc()의 "/hello" path에 대한 익명함수를 실행하여 "Hello World"를 출력한다.

- 익명함수 안의 http.ResponseWriter 파라미터는 HTTP Response에 무언가 쓸 수 있게하고, http.Request 파라미터는 입력된 Request요청을 검토할 수 있게 한다.
- ListenAndServe()메서드는 2개의 파라미터를 가지고 있는데 첫 번째는 포트 5000에 Request를 Listen할 것을 지정한다.
- 두 번째는 어떤 ServeMux를 사용할 지를 지정하는데 nil인 경우 DefaultServeMux를 사용한다.
- ServeMux는 기본적으로 HTTP Reqeuset Router(혹은 Multiplexor)인데 일반적으로 내장된 DefaultServeMux를 사용하지만 개발자가 별도로 ServeMux를 만들어 Routing을 관리할 수 있다.

- ### http.Handle()사용 예시
  - HandleFunc()과 비슷하게 HTTP Handler를 정의하는 다른 방식으로 http.Handle() 메서드를 사용할 수 있다.
  - 첫 번째 파라미터는 URL을 받고, 두 번째 파라미터로 http.Handle 인터페이스를 갖는 객체를 받아들인다.
  - http.Handle 인터페이스는 아래와 같은 메서드 하나를 갖는 인터페이스이다.
  ```go
  type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
  }
  ```
  - 아래 예제는 http.Handler 인터페이스를 갖는 testHandler 라는 struct를 정의하고 이 struct의 메서드 ServeHTTP()를 구현한 예시이다.
  - Handle()의 두 번째 파라미터는 testHandler 객체를 new()함수로 생성하여 전달한다.
  ```go
  package main
  
  import "net/http"
  
  func main() {
    http.Handle("/", new(testHandler))
    
    http.ListenAndServe(":5000", nil)
  }
  
  type testHandler struct {
    http.Handler
  }
  
  func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    str := "Hello World" + req.URL.Path
    w.Write([]byte(str))
  }
  ```
