package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // criei um slice com dois tamanhos
	s2 := append(s1, 1, 2, 3) // adicionei um valor no slice s1
	fmt.Println(s1, s2)

	s1[0] = 7           // alterei o espaço 0 para o elemento 7
	fmt.Println(s1, s2) // assim eu consigo provar eu consigo mudar as duas slice apartir de um.

}
