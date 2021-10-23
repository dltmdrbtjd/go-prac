package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

var users = map[string]*User{}

// 4번
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		// 들어오는 요청의 response header에 content-type을 추가한다.
		next.ServeHTTP(rw, r)
		// 전달 받은 http.handler를 호출한다.
	})
}

func main() {
	// 1.새로운 mux생성하기
	mux := http.NewServeMux()

	// 2.원래는 HandleFunc으로 만든 함수를 HandlerFunc 타입으로 변경
	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost:
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			users[user.Email] = &user
			json.NewEncoder(wr).Encode(user)
		}
	})
	// 3.만들어놓은 미들웨어에 userHandler를 파라미터로 넘긴다.
	// 그럼 미들웨어 함수에서 header에 content-type:application/json을 추가하고
	// userHandler를 호출하여 실행하게되는것 같음
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":5000", mux)
}

// 참고한 블로그
// https://woony-sik.tistory.com/12
