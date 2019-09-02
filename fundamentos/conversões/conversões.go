package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// concatenar uma int para uma string CUIDADO

	// isso vai dar errado, pois ele vai mostrar o numero da tabela ask que no caso é o 'a'
	fmt.Println("test" + string(101))
	// agora como deve ser feito a concatenação !!!
	fmt.Println("test" + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// boolean
	b, _ := strconv.ParseBool("True")
	if b {
		fmt.Println("Verdadeiro")
	}
} 
