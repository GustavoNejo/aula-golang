// Mais alguns exeplos de sites (gobyexemple.com) apartes do curso Udemy
package main

import "fmt"

func gvmain() {
	//cria um map com varios elementos
	m := make(map[string]int)
	m["k1"] = 19
	m["k2"] = 30
	fmt.Println("map", m)

	// pega um elemeto de um map
	v1 := m["k1"]
	fmt.Println("v1", v1)
	fmt.Println("len:", len(m))

	// deleta um elemento de um map
	delete(m, "k2")
	fmt.Println("map:", m)
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	//cria um map com o nome de cada um na frente de seu valor
	n := map[string]int{"fool": 1, "bar": 2}
	fmt.Println("map", n)
}
