// Mais alguns exeplos de sites (gobyexemple.com) apartes do curso Udemy

package main

import "fmt"

func bvkJmain() {
	s := make([]string, 4)
	fmt.Println("Emp:", s)

	s[0] = "g"
	s[1] = "u"
	s[2] = "s"
	s[3] = "t"

	fmt.Println("mostra:", s)  //todos os valores
	fmt.Println("pega:", s[2]) // pega o valor que eu quiser

	fmt.Println("conta:", len(s)) // conta quantos valores tem na linha vertical
	//len retorna o tamanho da fatia como esperado.

	// vai adicionar mais esse valor no slice s
	s = append(s, "a")
	s = append(s, "v", "o")
	fmt.Println("apd", s)

	fmt.Println("pegue os valores:", s[2], s[4], s[6], s[1])

	// fazendo uma copia
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // vai pegar os caracteres que esta entre os valores colocados
	fmt.Println("sl1:", l)

	l = s[:5] // para apagar o valor que estiver na casa 5 do slice
	fmt.Println("sl2", l)

	l = s[2:] // assim ele vai apagar todas as 'casa' que tiver para tras
	fmt.Println("sl3:", l)

	// iniciando variavel e slice apartir de uma inica linha
	t := []string{"m", "o", "r", "a", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
