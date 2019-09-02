package main

import (
	"fmt"
)

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2 
	comprarTv32 := trab1 != trab2 // isso é quivalente ao ou exclusivo (que nao existe em go)
	comprarSorvete := trab1 || trab2 // esses dois || significa ou

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true) // esses dois true significa que eu trabalhei os dois dias e consegui comprar a tv e sair para tomar sorverte
	fmt.Printf("tv50: %t, tv32: %t, Sorvete: %t, saudavel %t",
		tv50, tv32, sorvete, !sorvete) // Esse !sorvete é a negaçao logica
}
