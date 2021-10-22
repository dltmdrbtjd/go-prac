package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, t int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
	out.Flush()
}
