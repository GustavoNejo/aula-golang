package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preço float64
}

// interface é implementadas implicitamente !

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preço)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
func main() {
	var coisa imprimivel = pessoa{"Roberto", "Carlos"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça jeans", 89.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 120.20}
	fmt.Println(p2.toString())
	imprimir(p2)
}
