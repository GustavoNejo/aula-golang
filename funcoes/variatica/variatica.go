package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros { // Eu só não entendi o pq que tem esse _
		total += num
	}
	return total / float64(len(numeros)) //
}
func dsmain() {
	fmt.Printf("Média: %.2f", media(32, 9, 8))
}
