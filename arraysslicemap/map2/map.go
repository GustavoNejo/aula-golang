package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ //Criei e iniciei a funcESalario
		// adicionei alguns funcionarios
		"José João": 1133.44,
		"chaves":    5432.75,
		"Professor": 1334.65,
	}
	fmt.Println(funcsESalarios)
	funcsESalarios["Tesouro"] = 900.0   // aqui eu adicionei mais um funcionario
	delete(funcsESalarios, "José João") // deletei o funcionario Jose João

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario) // aqui automaticamente ele ja entende a onde que é o nome pq esta entre aspas e o salario esta fora das chaves
	}

}
