package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovados := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovados)
	}
}

func main() {
	aprovados := []string{"Maria", "rafael", "Pedro", "Guilherme"} //isso nao pode ser trata como uma array, pois uma array aqui daria um erro
	imprimirAprovados(aprovados...) // esse 'imprimirAprovados veio da func de acima 
}
