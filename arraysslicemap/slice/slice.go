package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // isso é uma array
	s2 := a2[1:3]               // isso é um slice que aponta para a array a2.
	//OBS slice não é uma array, um slice pega um pedaço da array para poder usar
	fmt.Println(a2, s2)

	s3 := a2[:2] // isso é um mesmo slice, que aponta para uma mesma array no caso a a2
	// esse slice vai pegar a quantidade de elento inserido apartir do elemento 0 da array. EX:
	fmt.Println(a2, s3)

	// tbm podemos fazer um slice de um slice. Ele vai pegar o s2 vai ver o primeiro elemento que foi imprido do primeiro slice.
	// OBS tds so slice criado foi gerado apartir do array no caso a a2
	s4 := s2[:1]
	fmt.Println(s2, s4)
}

// Mais alguns exeplos de sites apartes do curso Udemy
