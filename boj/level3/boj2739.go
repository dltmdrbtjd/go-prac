package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	for i := 1; i < 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}
