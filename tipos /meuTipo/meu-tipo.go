package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	}
	if n.entre(7.0, 8.99) {
		return "B"
	}
	if n.entre(5.0, 6.99) {
		return "C"
	}
	if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}
func amain() {
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(6.7))
	fmt.Println(notaParaConceito(1.5))
}
