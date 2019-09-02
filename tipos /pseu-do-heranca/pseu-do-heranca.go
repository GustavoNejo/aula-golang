package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro       // campo anonimo
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidade = 9
	f.turboLigado = false

	fmt.Printf("A ferrari %s está com turbo lidago ? %v\n", f.nome, f.turboLigado)
}
