package main

import (
	"fmt"
	"time"
)

func main() {
	// esse for vai contar de 1 até 10 de ordem crescente
	i := 1
	for i <= 10 { // nao tem while em go
		fmt.Println(i)
		i++
	}

	// esse for irá contar de 1 até 20 pulando de dois em dois
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}
	fmt.Println("\nMisturando ...")

	// esse for serve para mostrar impar e par.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}
	for {
		// quando eu abro a chave e nao escrevo nd, ele automaticamente vira um for infinito!!
		fmt.Println("Para sempre ...")
		time.Sleep(time.Second) // serve para imprimir algo de 1 em 1 segundo.
		// quer fazer com mais tempo é só multiplica pela quantidade que deseja. EX:
		//	time.Sleep(time.Second * 5)

	}
}
