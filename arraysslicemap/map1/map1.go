package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[123456789012] = "José"
	aprovados[982173988741] = "Pamela"
	aprovados[128779847022] = "Marcos"
	aprovados[238496791378] = "Cirillo"
	aprovados[746183279854] = "Cleice"
	aprovados[647389203817] = "Lucimar"
	fmt.Println(aprovados)

	for CPF, nome := range aprovados {
		fmt.Printf("%s (CPF %d)\n", nome, CPF)
	}

	//fmt.Printf(aprovados[123456789012])
	delete(aprovados, 238496791378)
	fmt.Println(aprovados[238496791378]) // esse Print serve para mostrar que foi deletado esse cpf
}
