package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int                 // aqui o slice esta vazio
	slice1 = append(slice1, 4, 5, 6) // aqui no slice foi adicionado alguns elementos (4,5,6) no caso
	// array1 = append(array1, 4,5,6) isso nao é possivel pq uma array nao consegue receber outra array
	fmt.Println(array1, slice1)

	// copy simplismente vai pegar o que esta escrito em um slice e passar para o outro

	slice2 := make([]int, 2) // neste caso ele vai pegar os dois primeiros elementos do slice um e passar para o slice2
	copy(slice2, slice1)     // lembrando que o primeiro slice é sempre o que vai esta recebendo a copia
	// eu nao posso possar para copy diretamente um array! Tem sempre que ser um copy
	fmt.Println(slice2)

}
