package main

import (
	"encoding/json"
	"fmt"
)

// criando um Json
// coloque primeiro o que voce quer passar (Ex "ID, Nome") a primeira letra em maiusculo
// criando um json com letra maiuscula ele sara PUBLICO e com letra minuscula ele sera PRIVADO
// Diferença entra publico e privado é que se voce coloca isso dentro de um pacote e ele é privado ele nao sera visto fora do pacote, se ele for publico ele sera visto em ambos

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preço float64  `json:"preço"`
	Tags  []string `json:"tags"`
}

// struct para um json
func main() {
	p1 := produto{1, "Notebook", 1899.90, []string{"Promocao", "Eletronicos"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Json para struct
	var p2 produto
	jsonString := `{"id":2 , "nome": "caneta", "preco":8.90,"Tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
