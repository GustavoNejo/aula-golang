package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Print("liretal inteiro é:", reflect.TypeOf(2134))

	// sem sinal (só dispositovos)... uint08, uint16, uint32, uint64
	var b byte = 12
	fmt.Println("o byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int 64
	i1 := math.MaxInt64
	fmt.Println("O maior valor de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	// representa o mapeamento da tabela unicode (int32)
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32 e float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))              //float32
	fmt.Println("O Tipo literal 49.99 é:", reflect.TypeOf(49.99)) // float64

	// tipo boolean
	bo := true
	fmt.Println("O tipo de bo é um:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "meu nome é Gustavo"
	fmt.Println(s1)                                // essa '!' é inserido de acordo com a concatenação(+).
	fmt.Println("O tamanho da string é:", len(s1)) // esse 'len' serve para contar quantos caracteres tem na variavel

	//string com multiplas linhas
	s2 := `Olá
meu
nome
é
Gustavo`
	fmt.Println("O tamanho da string é ``:", len(s2))
}
