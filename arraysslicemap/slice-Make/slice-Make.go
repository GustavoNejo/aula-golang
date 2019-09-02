package main

import "fmt"

func main() {
	// criando um slice
	s := make([]int, 10) // criei uma variavel array do nome make      //criei um slice de tem espaço de 10 elementos
	s[9] = 12            // o mome do slice s          // no nono espaço (ultimo) eu adicionei o elemento 12
	fmt.Println(s)

	s = make([]int, 10, 20)                     // ele crio um segundo slice apontado para array make, só que aqui ele vai partir ultima quantidade de elementos do array, ele vai meio que reconfigurar o slice. Ex se tnha 10 espaço e vc faz isso ele vai par 20
	fmt.Println(s, len(s), cap(s))              // len = tamanho, cap = capacidade
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // func append serve para adicionar em um slice que ja esta criado(mais ou menos como uma var)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // fazer isso o slice vai dobrar de tanho, caso já tenha atingido o tamanho maximo dele (ele nao vai dar erro e nem nd)
	fmt.Println(s, len(s), cap(s))
}
