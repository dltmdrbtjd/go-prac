package main

import "fmt"

func main() {
	var n, s int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		s += i
	}
	fmt.Println(s)
}
