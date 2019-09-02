package main

import (
	"fmt"
	"time"
)

func main() {
	// Aqui ele vai fazer uma comparaçao para ver se os nomes sao iguais caso for ele vai retornar o valor true ou false

	fmt.Println("String:", "banana" == "banana")
	fmt.Println("!=", 3 != 2) //!= diferente
	fmt.Println("<", 3 < 2)   // < menor
	fmt.Println(">", 3 > 2)   // > maior
	fmt.Println("<=", 3 <= 2) // <= menor ou igual
	fmt.Println(">=", 3 >= 2) // >= maior ou igual

	// trabalhando com datas
	d1 := time.Unix(0, 0) // esse time.Unix serve para data
	d2 := time.Unix(0, 0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))
}
