package main

import "fmt"

type produto struct {
	// OBS para declarar eu nao uso virgulas
	nome     string
	preco    float64
	desconto float64
}

// Metodo: funcao com receiver (receptor )

func (p produto) precoComDesconto() float64 { // o nome da minha func vai ser produto, mas que está referenciado para p
	return p.preco * (1 - p.desconto) // nesse caso eu uso p.preco para referenciar o preço que esta no struct acima.
	//para fazer um desconto eu uso essa formula da linha 12
}

func main() {
	var produto1 produto
	produto1 = produto{
		// lembrando no final de cada linha aqui eu tenho que adicionar uma vergula !!
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto()) // dessa forma ele vai pegar o valor do produto e o valor do desconto e fazer o calculo do desconto !
	// um segundo modo de fazer um struct mais simples é assim:
	produto2 := produto{"Notebook", 1989.90, 0.10} // assim eu posso fazer um struct simplificado contendo as mesma informações.
	// OBS eu tenho que criar o segundo da mesma forma que eu criei o type struct. 
	fmt.Println(produto2, produto2.precoComDesconto())

}
