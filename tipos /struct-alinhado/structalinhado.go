package main

import "fmt"

// Primeiro struct serve para armazenar as config dos itens
type item struct {
	produtoID int
	qtde      int
	preco     float64
}

// Segunda struct serve para armazenar as config dos peidos
type pedido struct {
	userID int
	itens  []item // Lembrando que o colchetes sem nd é pq é um slice
}

// Esse bloco vai ser feita o calculo dos pedidos
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

// Esse bloco vai mostrar quais foram os pedido feitos, a quantidade e o preço deles
func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10}, // o struct pode ser feito assim, de um forma mais clara porem eu vou digitar mais
			item{2, 1, 23.49},                         // ou assim uma forma menos clara, porem eu vou digitar memos
			item{produtoID: 11, qtde: 200, preco: 3.13},
		}, // nao posso esquecer dessa virguala !!
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal()) // Esse pedido.valorTotal significa que ele vai na func pedidos vai buscar o Valor Total e me devolver um float64
}
