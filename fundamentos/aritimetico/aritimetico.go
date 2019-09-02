package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma =", a+b)
	fmt.Println("subtração", a-b)
	fmt.Println("divisao =", a/b)
	fmt.Println("multiplicaçao  =>", a*b)
	fmt.Println("modulo =>", a%b)

	// bitwise
	fmt.Println("E => ", a&b)
	fmt.Println("Ou =>", a|b)

	c := 3.0
	d := 2.0

	// outras variaves usando o meth
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // math.Max serve para mostrar o maior valor entre as variaveis
	fmt.Println("Menor =>", math.Min(c, d)) //math.Min vai mostrar o menor numero entre as variaveis
	fmt.Println("Exponenciação =>", math.Pow(c, d)) //math.Pow vai motrar a exponenciancao entre as variaveis
	 
}
