package main

import (
	"net/http"

	"api-prac/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHadnler())
}
