package main

import "fmt"

func main() {
	i := 1

	// go nao tem aritimetica de ponteiros

	// como criar um ponteiro em go ?
	var p *int = nil // nil é igul a nd/vazil

	p = &i // pegando o endereço da variavel i
	*p++   // desreferenciando (pegando somente o valor da variavel acrecentando um valor)
	i++
	fmt.Println(*p, i, &i)

}
