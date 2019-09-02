package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(função func(int, int) int, p1, p2 int) int {
	return função(p1, p2)
}
func main() {
	resultado := exec(multiplicacao, 3, 5)
	fmt.Println(multiplicacao(3, 5))
	fmt.Println(resultado)
}
 