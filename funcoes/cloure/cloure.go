package main

import "fmt"

func cloure() func() { // aqui eu tenho uma func que vai retornar outra func
	x := 10
	var funcaoExistente = func() { // aqui eu criei uma var que se chama funcaoExistente e atribui para uma func
		fmt.Println(x)
	}
	return funcaoExistente // que no caso é uma func
}
func main() {
	x := 25
	fmt.Println(x)
	imprimeX := cloure()
	imprimeX()
}
