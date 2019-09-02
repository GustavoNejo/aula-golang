package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!") // esse defer serve para fechar um recurso
	if numero > 5000 {
		fmt.Println("valor alto...")
		return numero
	} // nessa linha aqui tambem poderia haver um else, mas nao tem necessidade
	fmt.Println("valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(2133))
}
