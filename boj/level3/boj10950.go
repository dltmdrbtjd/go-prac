package main

import "fmt"

func main() {
	var a int
	var b int
	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}

}
