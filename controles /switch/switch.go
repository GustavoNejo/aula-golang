package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // serve para o code continuar descendo para os outros cases. Sempre que quiser que ele execulte os case de baixoa
	case 9:
		return "A"

		// pode ser execultado dos dois modos, tanto do jeito do case 9 como do modo do case 8 e 7
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}
func main() {
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(5.5))
	fmt.Println(notaParaConceito(4.9))
	fmt.Println(notaParaConceito(0))
	fmt.Println(notaParaConceito(-1))
	//fmt.Println(notaParaConceito(sete))
}
