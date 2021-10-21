package main

import "fmt"

func main() {
	var h, m int

	fmt.Scan(&h, &m)

	m -= 45
	if m < 0 {
		m += 60
		h -= 1
	}

	if h < 0 {
		h += 24
	}

	fmt.Println(h, m)
}
