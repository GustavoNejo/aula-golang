package main

import "fmt"

func notasMaps(nota int) {
	var array = []string{10: "A", 9: "A", 8: "B", 7: "B", 5: "C", 4.: "C", 3: "D", 2: "D", 1: "E", 0: "E"}
	if nota > 10 {
		fmt.Println("Nota invalida !")
		return
	}
	if nota < 0 {
		fmt.Println("Nota invalida !")
		return
	}
	resultado := array[nota]
	fmt.Println(resultado)
	return
}
func main() {
	notasMaps(9)
	notasMaps(10)
	notasMaps(10)

}
