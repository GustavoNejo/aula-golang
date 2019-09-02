package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador de contas
	for i, numero := range numeros {
		fmt.Printf("%d) (%d)\n", i, numero)
	}
	// eu tambem posso usar ignorar o valor do indice se eu quiser. Podemos tbm tirar o _ mas assim eu só vou acessar o indice da array.

	for _, num := range numeros {
		fmt.Println(num)
	}
}
