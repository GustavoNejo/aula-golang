package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, error := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-1)
	if err != nil {
		fmt.Println(err, error)
	}
}
