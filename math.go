package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mult(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	return a / b
}
