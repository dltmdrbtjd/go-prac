package main

func main() {
	var a interface{} = 1

	i := a
	j := a.(int)

	println(i)
	println(j)
}
