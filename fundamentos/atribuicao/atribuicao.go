package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 33 // inferencia
	fmt.Println(i)
	i += 3 //i = i = 3
	i -= 3 //i = 1 - 3
	i /= 2 //i = i / 2
	i *= 2 //i = i * 2
	i %= 2 //i = i * 2
	fmt.Println(i)

	x, y := 1, 2
	x, y = y, x  // Isso serve para poder trocar os valores (o valor de x vai y e o y vai para o x)
	
}
