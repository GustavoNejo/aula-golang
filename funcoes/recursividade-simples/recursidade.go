package main

import "fmt"

func fatorial(n uint) uint { // uint signifca que eu vou pegar um valor inteiro que nao tem sinal(nao pode ser um valor negativo)
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(2)
	fmt.Println(resultado)
}
