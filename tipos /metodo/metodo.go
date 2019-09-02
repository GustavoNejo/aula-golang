package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string { // aqui ele fez p para acessar as pessoas.
	return p.nome + " " + p.sobrenome //aqui vai retorna o preimeiro e segundo nome, para no final juntar os dois
}
func (p *pessoa) setNomeCompleto(nomeCompleto string) { // Aqui eu criei um ponteiro para pessoa
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]

}
func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("João Pedro")
	fmt.Println(p1.getNomeCompleto())

}
