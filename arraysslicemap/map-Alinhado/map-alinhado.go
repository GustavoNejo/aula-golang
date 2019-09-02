package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Amanda Silva": 1200.0,
			"Antonio":      1600.60,
		},
		"B": {
			"Beatriz": 3000.00,
			"Bianca":  1800.0,
		},
		"C": {
			"Carlos": 900.00,
			"Camila": 4000.98,
		},
		"D": {
			"Daniel": 4983.13,
		},
	}
	delete(funcsPorLetra, "D") // OBS nao tem como apagar somente um elemento da chave, eu só consigo deletar a chave inteira !!!
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
