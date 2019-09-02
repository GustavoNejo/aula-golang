package main

import "fmt"

func main() {
	// um array é uma estrutura estatica (fixo) e homogenia (mesmo tipo).
	// exemplo se voce cria uma array de inteiros só vai poder conter inteiros dentro dele
	// é uma strutura estatica, não cresce de tamanho
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.2
	// Lembrando que os arrays é começado a contar apartir do indenci que é o mesmo que 0
	// Entao se eu tentar colocar mais uma 'nota' ele vai dar um erro de compilaçao. EX:
	// notas[3] = 10
	fmt.Print(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	} 

	media := total / float64(len(notas))
	fmt.Printf("\nMédia %.2f\n", media)
}
