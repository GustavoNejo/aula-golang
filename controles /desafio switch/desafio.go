package main

import "fmt"
// isso é uma refatoração do if-else-if para switch

func notaParaConceito(n float64) string {
	switch true{
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.0))
	fmt.Println(notaParaConceito(8.0))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.0))

}
